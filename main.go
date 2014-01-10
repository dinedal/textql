package main

import (
    "database/sql"
    "encoding/csv"
    "flag"
    "fmt"

    _ "github.com/mattn/go-sqlite3"

    "bytes"
    "io"
    "log"
    "os"
    "os/user"
    "path/filepath"
    "strconv"
    "strings"
    "time"
    "unicode/utf8"
)

func main() {
    // Open in memory db
    db, _ := sql.Open("sqlite3", ":memory:")
    defer db.Close()

    // Parse command line opts
    commands := flag.String("sql", "", "SQL Command(s) to run on the data")
    source_text := flag.String("source", "stdin", "Source file to load, or defaults to stdin")
    delimiter := flag.String("dlm", ",", "Delimiter between fields, (\\t for tab)")
    header := flag.Bool("header", false, "Treat file as having the first row as a header row")
    flag.Parse()

    // Open the input source (don't close stdin)
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
        first_row = []string{}
    } else {
        headerRow = make([]string, len(first_row))
        for i := 0; i < len(first_row); i++ {
            headerRow[i] = "c" + strconv.Itoa(i)
        }
    }

    tableName := "tbl"

    createTable(&tableName, &headerRow, db)
    t0 := time.Now()
    tx, _ := db.Begin()
    //Load first row
    loadRow(&tableName, &first_row, tx)
    // Read the data
    eof := false

    for eof == false {
        row, file_err := reader.Read()
        if file_err == io.EOF {
            eof = true
        } else {
            loadRow(&tableName, &row, tx)
        }
    }
    tx.Commit()
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
    fmt.Println(buffer.String())
    _, err := db.Exec(buffer.String())
    if err != nil {
        log.Fatal(err)
    }
    return err
}

func loadRow(tableName *string, values *[]string, db *sql.Tx) error {
    if len(*values) == 0 {
        return nil
    }
    var buffer bytes.Buffer
    vals := make([]interface{}, 0)
    buffer.WriteString("INSERT INTO " + (*tableName) + " VALUES (")
    for i, val := range *values {
        buffer.WriteString("?")
        if i != len(*values)-1 {
            buffer.WriteString(", ")
        }
        vals = append(vals, val)
    }
    buffer.WriteString(");")
    //    fmt.Println(buffer.String())
    stmt, _ := db.Prepare(buffer.String())
    _, err := stmt.Exec(vals...)
    if err != nil {
        log.Fatal(err)
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
