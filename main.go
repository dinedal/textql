package main

import (
	"bytes"
	"database/sql"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/alecthomas/kingpin.v1"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	app          = kingpin.New(APP_NAME, APP_DESCRIPTION)
	queries      = app.Flag(FLAG_QUERIES, FLAG_QUERIES_HELP).String()
	input        = app.Flag(FLAG_INPUT, FLAG_INPUT_HELP).Default("stdin").String()
	delimiter    = app.Flag(FLAG_DELIMITER, FLAG_DELIMITER_HELP).Default(",").String()
	lazyQuotes   = app.Flag(FLAG_LAZY_QUOTES, FLAG_LAZY_QUOTES_HELP).Bool()
	header       = app.Flag(FLAG_HEADER, FLAG_HEADER_HELP).Bool()
	outputHeader = app.Flag(FLAG_OUTPUT_HEADER, FLAG_OUTPUT_HEADER_HELP).Bool()
	tableName    = app.Flag(FLAG_TABLE_NAME, FLAG_TABLE_NAME_HELP).Default(DEFAULT_TABLE_NAME).String()
	saveTo       = app.Flag(FLAG_SAVE_TO, FLAG_SAVE_TO_HELP).String()
	console      = app.Flag(FLAG_CONSOLE, FLAG_CONSOLE_HELP).Bool()
	verbose      = app.Flag(FLAG_VERBOSE, FLAG_VERBOSE_HELP).Bool()
)

func main() {
	app.Parse(os.Args[1:])

	if *console && *input == "stdin" {
		log.Fatalln(ERR_NO_CONSOLE_STDIN)
	}

	separator := determineSeparator(delimiter)

	// Open db, in memory if possible
	db, openPath := openDB(saveTo, console)

	// Open the input source
	var fp *os.File
	fp = openFileOrStdin(input)
	defer fp.Close()

	// Init a structured text reader
	reader := csv.NewReader(fp)
	reader.FieldsPerRecord = 0
	reader.Comma = separator
	reader.LazyQuotes = *lazyQuotes

	// Read the first row
	first_row, read_err := reader.Read()

	if read_err != nil {
		log.Fatalln(read_err)
	}

	var headerRow []string

	if *header {
		headerRow = first_row
		first_row, read_err = reader.Read()

		if read_err != nil {
			log.Fatalln(read_err)
		}
	} else {
		headerRow = make([]string, len(first_row))
		for i := 0; i < len(first_row); i++ {
			headerRow[i] = "c" + strconv.Itoa(i)
		}
	}

	// Create the table to load to
	createTable(tableName, &headerRow, db, verbose)

	// Start the clock for importing
	t0 := time.Now()

	// Create transaction
	tx, tx_err := db.Begin()

	if tx_err != nil {
		log.Fatalln(tx_err)
	}

	// Load first row
	stmt := createLoadStmt(tableName, &headerRow, tx)
	loadRow(tableName, &first_row, tx, stmt, verbose)

	// Read the data
	for {
		row, file_err := reader.Read()
		if file_err == io.EOF {
			break
		} else if file_err != nil {
			log.Println(file_err)
		} else {
			loadRow(tableName, &row, tx, stmt, verbose)
		}
	}
	stmt.Close()
	tx.Commit()

	t1 := time.Now()

	if *verbose {
		fmt.Fprintf(os.Stderr, "Data loaded in: %v\n", t1.Sub(t0))
	}

	// Determine what sql to execute
	sqls_to_execute := strings.Split(*queries, ";")

	t0 = time.Now()

	// Execute given SQL
	for _, sql_cmd := range sqls_to_execute {
		if strings.Trim(sql_cmd, " ") != "" {
			result, err := db.Query(sql_cmd)
			if err != nil {
				log.Fatalln(err)
			}
			displayResult(result, outputHeader, separator)
		}
	}

	t1 = time.Now()

	if *verbose {
		fmt.Fprintf(os.Stderr, "Queries run in: %v\n", t1.Sub(t0))
	}

	// Open console
	if *console {
		db.Close()
		args := []string{*openPath}
		if *outputHeader {
			args = append(args, "-header")
		}
		cmd := exec.Command("sqlite3", args...)

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		FLAG_err := cmd.Run()
		if cmd.Process != nil {
			cmd.Process.Release()
		}

		if len(*saveTo) == 0 {
			os.RemoveAll(filepath.Dir(*openPath))
		}

		if FLAG_err != nil {
			log.Fatalln(FLAG_err)
		}
	} else if len(*saveTo) == 0 {
		db.Close()
		os.Remove(*openPath)
	} else {
		db.Close()
	}
}

func createTable(tableName *string, columnNames *[]string, db *sql.DB, verbose *bool) error {
	var buffer bytes.Buffer

	buffer.WriteString("CREATE TABLE IF NOT EXISTS " + (*tableName) + " (")

	for i, col := range *columnNames {
		var col_name string

		reg := regexp.MustCompile(`[^a-zA-Z0-9]`)

		col_name = reg.ReplaceAllString(col, "_")
		if *verbose && col_name != col {
			fmt.Fprintf(os.Stderr, "Column %x renamed to %s\n", col, col_name)
		}

		buffer.WriteString(col_name + " TEXT")

		if i != len(*columnNames)-1 {
			buffer.WriteString(", ")
		}
	}

	buffer.WriteString(");")

	_, err := db.Exec(buffer.String())

	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func createLoadStmt(tableName *string, values *[]string, db *sql.Tx) *sql.Stmt {
	if len(*values) == 0 {
		log.Fatalln("Nothing to build insert with!")
	}
	var buffer bytes.Buffer

	buffer.WriteString("INSERT INTO " + (*tableName) + " VALUES (")
	for i := range *values {
		buffer.WriteString("?")
		if i != len(*values)-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString(");")
	stmt, err := db.Prepare(buffer.String())
	if err != nil {
		log.Fatalln(err)
	}
	return stmt
}

func loadRow(tableName *string, values *[]string, db *sql.Tx, stmt *sql.Stmt, verbose *bool) error {
	if len(*values) == 0 {
		return nil
	}
	vals := make([]interface{}, 0)
	for _, val := range *values {
		vals = append(vals, val)
	}
	_, err := stmt.Exec(vals...)
	if err != nil && *verbose {
		fmt.Fprintln(os.Stderr, "Bad row: ", err)
	}
	return err
}

type csvWriter struct {
	*csv.Writer
}

func (w csvWriter) put(record []string) {
	if err := w.Write(record); err != nil {
		log.Fatalln(err)
	}
}

func displayResult(rows *sql.Rows, outputHeader *bool, sep rune) {
	cols, cols_err := rows.Columns()

	if cols_err != nil {
		log.Fatalln(cols_err)
	}

	out := csvWriter{csv.NewWriter(os.Stdout)}

	out.Comma = sep

	if *outputHeader {
		out.put(cols)
	}

	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))

	dest := make([]interface{}, len(cols))
	for i := range cols {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		rows.Scan(dest...)

		for i, raw := range rawResult {
			result[i] = string(raw)
		}

		out.put(result)
	}

	out.Flush()
}

func openFileOrStdin(path *string) *os.File {
	var fp *os.File
	var err error
	if (*path) == "stdin" {
		fp = os.Stdin
		err = nil
	} else {
		fp, err = os.Open(*cleanPath(path))
	}

	if err != nil {
		log.Fatalln(err)
	}

	return fp
}

func cleanPath(path *string) *string {
	var result string
	usr, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}

	if (*path)[:2] == "~/" {
		dir := usr.HomeDir + "/"
		result = strings.Replace(*path, "~/", dir, 1)
	} else {
		result = (*path)
	}

	abs_result, abs_err := filepath.Abs(result)
	if abs_err != nil {
		log.Fatalln(err)
	}

	clean_result := filepath.Clean(abs_result)

	return &clean_result
}

func openDB(path *string, no_memory *bool) (*sql.DB, *string) {
	openPath := ":memory:"
	if len(*path) != 0 {
		openPath = *cleanPath(path)
	} else if *no_memory {
		outDir, err := ioutil.TempDir(os.TempDir(), "textql")
		if err != nil {
			log.Fatalln(err)
		}
		openPath = filepath.Join(outDir, "textql.db")
	}

	db, err := sql.Open("sqlite3", openPath)

	if err != nil {
		log.Fatalln(err)
	}
	return db, &openPath
}

func determineSeparator(delimiter *string) rune {
	var separator rune

	if (*delimiter) == "tab" {
		separator = '\t'
	} else if strings.Index((*delimiter), "0x") == 0 {
		dlm, hex_err := hex.DecodeString((*delimiter)[2:])

		if hex_err != nil {
			log.Fatalln(hex_err)
		}

		separator, _ = utf8.DecodeRuneInString(string(dlm))
	} else {
		separator, _ = utf8.DecodeRuneInString(*delimiter)
	}
	return separator
}
