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
    flag.Parse()

    if *console && (*source_text == "stdin") {
        fmt.Println("Can not open console with pipe input, read a file instead")
        os.Exit(-1)
    }

    // Open in memory db
    db, openPath := openDB(save_to, console)

    // Open the input source
    var fp *os.File
    fp = openFileOrStdin(source_text)
    defer fp.Close()

    // Init a structured text reader
    reader := csv.NewReader(fp)
    reader.FieldsPerRecord = 0

    // Define the seperator
    var seperator rune

    if (*delimiter) == "\\t" {
        seperator = '\t'
    } else {
        seperator, _ = utf8.DecodeRuneInString(*delimiter)
    }
    reader.Comma = seperator

    // Read the first row
    first_row, _ := reader.Read()
    var headerRow []string

    if *header {
        headerRow = first_row
        first_row, _ = reader.Read()
    } else {
        headerRow = make([]string, len(first_row))
        for i := 0; i < len(first_row); i++ {
            headerRow[i] = "c" + strconv.Itoa(i)
        }
    }

    createTable(tableName, &headerRow, db)
    t0 := time.Now()

    stmt := createLoadStmt(tableName, &headerRow, db)

    //Create transaction
    if *openPath == ":memory:" {
        //Load first row
        loadDBRow(tableName, &first_row, db, stmt)

        // Read the data
        eof := false

        for eof == false {
            row, file_err := reader.Read()
            if file_err == io.EOF {
                eof = true
            } else {
                loadDBRow(tableName, &row, db, stmt)
            }
        }
    } else {
        tx, _ := db.Begin()
        loadTXRow(tableName, &first_row, tx, stmt)

        // Read the data
        eof := false

        for eof == false {
            row, file_err := reader.Read()
            if file_err == io.EOF {
                eof = true
            } else {
                loadTXRow(tableName, &row, tx, stmt)
            }
        }
        tx.Commit()
    }

    t1 := time.Now()

    fmt.Printf("Data loaded in: %v\n", t1.Sub(t0))
    // Determine what sql to execute
    sqls_to_execute := strings.Split(*commands, ";")

    // Execute given SQL
    for _, sql_cmd := range sqls_to_execute {
        if strings.Trim(sql_cmd, " ") != "" {
            result, err := db.Query(sql_cmd)
            if err != nil {
                log.Fatal(err)
            }
            displayResult(result)
        }
    }

    // Open console
    if *console {
        db.Close()
        cmd := exec.Command("/usr/bin/sqlite3", *openPath)
        // cmd := exec.Command("/bin/bash")
        //        cmd := exec.Command("read")
        //cmd := exec.Command("date")
        fmt.Println(os.Stdin.Name())
        //        os.Stdin.Truncate(0)
        //        os.Stdin.Close()
        cmd.Stdin = os.Stdin
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        c_err := cmd.Run()
        fmt.Println(c_err)
        fmt.Println("HELLOOOOO   ", *openPath)
        if len(*save_to) == 0 {
            os.Remove(*openPath)
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
        log.Fatal(err)
    }
    return err
}

func createLoadStmt(tableName *string, values *[]string, db *sql.DB) *sql.Stmt {
    if len(*values) == 0 {
        log.Fatal("Nothing to build insert with!")
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
        log.Fatal(err)
    }
    return stmt
}

func loadDBRow(tableName *string, values *[]string, db *sql.DB, stmt *sql.Stmt) error {
    if len(*values) == 0 {
        return nil
    }
    vals := make([]interface{}, 0)
    for _, val := range *values {
        vals = append(vals, val)
    }
    _, err := stmt.Exec(vals...)
    if err != nil {
        log.Println("Bad row: ", err)
    }
    return err
}

func loadTXRow(tableName *string, values *[]string, db *sql.Tx, stmt *sql.Stmt) error {
    if len(*values) == 0 {
        return nil
    }
    vals := make([]interface{}, 0)
    for _, val := range *values {
        vals = append(vals, val)
    }
    _, err := stmt.Exec(vals...)
    if err != nil {
        log.Println("Bad row: ", err)
    }
    return err
}

func displayResult(rows *sql.Rows) {
    cols, _ := rows.Columns()
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
        log.Fatal(err)
    }

    return fp
}

func cleanPath(path *string) *string {
    var result string
    usr, err := user.Current()
    if err != nil {
        log.Fatal(err)
    }

    dir := usr.HomeDir + "/"
    if (*path)[:2] == "~/" {
        result = strings.Replace(*path, "~/", dir, 1)
    } else {
        result = (*path)
    }

    abs_result, abs_err := filepath.Abs(result)
    if abs_err != nil {
        log.Fatal(err)
    }
    return &abs_result
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
        log.Fatal(err)
    }
    return db, &openPath
}
