package main

import (
    "database/sql"
    "encoding/csv"
    "flag"
    "fmt"

    _ "github.com/mattn/go-sqlite3"

    "bytes"
    "crypto/rand"
    "io"
    "log"
    "os"
    "os/exec"
    "os/user"
    "path/filepath"
    "strconv"
    "strings"
    "time"
    "unicode/utf8"
)

func main() {
    // Parse command line opts
    commands := flag.String("sql", "", "SQL Command(s) to run on the data")
    source_text := flag.String("source", "stdin", "Source file to load, or defaults to stdin")
    delimiter := flag.String("dlm", ",", "Delimiter between fields, (\\t for tab)")
    header := flag.Bool("header", false, "Treat file as having the first row as a header row")
    tableName := flag.String("table-name", "tbl", "Override the default table name (tbl)")
    save_to := flag.String("save-to", "", "If set, sqlite3 db is left on disk at this path")
    console := flag.Bool("console", false, "After all commands are run, open sqlit3 console with this data")
    verbose := flag.Bool("verbose", false, "Enable verbose logging")
    flag.Parse()

    if *console && (*source_text == "stdin") {
        log.Fatalln("Can not open console with pipe input, read a file instead")
    }

    seperator := determineSeperator(delimiter)

    // Open db, in memory if possible
    db, openPath := openDB(save_to, console)

    // Open the input source
    var fp *os.File
    fp = openFileOrStdin(source_text)
    defer fp.Close()

    // Init a structured text reader
    reader := csv.NewReader(fp)
    reader.FieldsPerRecord = 0
    reader.Comma = seperator

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
    createTable(tableName, &headerRow, db)

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
            log.Fatalln(file_err)
        } else {
            loadRow(tableName, &row, tx, stmt, verbose)
        }
    }
    tx.Commit()

    t1 := time.Now()

    if *verbose {
        fmt.Printf("Data loaded in: %v\n", t1.Sub(t0))
    }

    // Determine what sql to execute
    sqls_to_execute := strings.Split(*commands, ";")

    // Execute given SQL
    for _, sql_cmd := range sqls_to_execute {
        if strings.Trim(sql_cmd, " ") != "" {
            result, err := db.Query(sql_cmd)
            if err != nil {
                log.Fatalln(err)
            }
            displayResult(result)
        }
    }

    // Open console
    if *console {
        db.Close()

        cmd := exec.Command("sqlite3", *openPath)
        cmd.Stdin = os.Stdin
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        cmd_err := cmd.Run()

        if len(*save_to) == 0 {
            os.Remove(*openPath)
        }

        if cmd_err != nil {
            log.Fatalln(cmd_err)
        }
    } else if len(*save_to) == 0 {
        db.Close()
        os.Remove(*openPath)
    } else {
        db.Close()
    }
}

func createTable(tableName *string, columnNames *[]string, db *sql.DB) error {
    var buffer bytes.Buffer
    buffer.WriteString("CREATE TABLE " + (*tableName) + " (")
    for i, col := range *columnNames {
        buffer.WriteString(col + " TEXT")
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
        log.Println("Bad row: ", err)
    }
    return err
}

func displayResult(rows *sql.Rows) {
    cols, cols_err := rows.Columns()

    if cols_err != nil {
        log.Fatalln(cols_err)
    }

    rawResult := make([][]byte, len(cols))
    result := make([]string, len(cols))

    dest := make([]interface{}, len(cols))
    for i, _ := range cols {
        dest[i] = &rawResult[i]
    }

    for rows.Next() {
        rows.Scan(dest...)

        for i, raw := range rawResult {
            result[i] = string(raw)
        }

        for j, v := range result {
            fmt.Printf("%s", v)
            if j != len(result)-1 {
                fmt.Printf(", ")
            }
        }
        fmt.Printf("\n")
    }
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
        b := make([]byte, 10)
        io.ReadFull(rand.Reader, b)
        tmpPath := "/tmp/dankbase_" + fmt.Sprintf("%x", b) + ".db"
        openPath = tmpPath
    }

    db, err := sql.Open("sqlite3", openPath)

    if err != nil {
        log.Fatalln(err)
    }
    return db, &openPath
}

func determineSeperator(delimiter *string) rune {
    var seperator rune

    if (*delimiter) == "\\t" {
        seperator = '\t'
    } else {
        seperator, _ = utf8.DecodeRuneInString(*delimiter)
    }
    return seperator
}
