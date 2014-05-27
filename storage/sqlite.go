package storage

import (
	"bytes"
	"database/sql"
	"fmt"
	"strings"

	"github.com/dinedal/textql/inputs"
	sqlite3 "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"regexp"
)

type sqlite3Storage struct {
	options *SQLite3Options
	db      *sql.DB
	connId  int
}

type SQLite3Options struct {
	//	SaveToPath string
}

var (
	sqlite3conn []*sqlite3.SQLiteConn = []*sqlite3.SQLiteConn{}
)

func NewSQLite3Storage(opts *SQLite3Options) *sqlite3Storage {
	this := &sqlite3Storage{
		options: opts,
	}

	sql.Register("sqlite3_textql",
		&sqlite3.SQLiteDriver{
			ConnectHook: func(conn *sqlite3.SQLiteConn) error {
				sqlite3conn = append(sqlite3conn, conn)
				return nil
			},
		})

	this.open()
	return this
}

func (this *sqlite3Storage) open() {
	db, err := sql.Open("sqlite3_textql", ":memory:")

	if err != nil {
		log.Fatalln(err)
	}

	db.Ping()
	this.connId = len(sqlite3conn) - 1
	this.db = db
}

func (this *sqlite3Storage) LoadInput(input inputs.Input) {
	this.createTable("tbl", input.Header(), true)

	tx, tx_err := this.db.Begin()

	if tx_err != nil {
		log.Fatalln(tx_err)
	}

	stmt := this.createLoadStmt("tbl", len(input.Header()), tx)

	row := input.ReadRecord()
	for {
		if row == nil {
			break
		}
		this.loadRow("tbl", len(input.Header()), row, tx, stmt, true)
		row = input.ReadRecord()
	}
	stmt.Close()
	tx.Commit()
}

func (this *sqlite3Storage) createTable(tableName string, columnNames []string, verbose bool) error {
	var buffer bytes.Buffer

	tableNameCheckRegEx := regexp.MustCompile(`.*\[.*\].*`)

	if tableNameCheckRegEx.FindString(tableName) != "" {
		log.Fatalln("Invalid table name", tableName)
	}

	buffer.WriteString("CREATE TABLE IF NOT EXISTS [" + (tableName) + "] (")

	for i, col := range columnNames {
		columnNameCheckRegEx := regexp.MustCompile(`.*\[.*\].*`)

		if columnNameCheckRegEx.FindString(col) != "" {
			log.Fatalln("Invalid table name", col)
		}

		buffer.WriteString("[" + col + "] TEXT")

		if i != len(columnNames)-1 {
			buffer.WriteString(", ")
		}
	}

	buffer.WriteString(");")

	_, err := this.db.Exec(buffer.String())

	if err != nil {
		log.Fatalln(err)
	}

	if verbose {
		log.Println(buffer.String())
	}

	return err
}

func (this *sqlite3Storage) createLoadStmt(tableName string, colCount int, db *sql.Tx) *sql.Stmt {
	if colCount == 0 {
		log.Fatalln("Nothing to build insert with!")
	}
	var buffer bytes.Buffer

	buffer.WriteString("INSERT INTO " + (tableName) + " VALUES (")
	// Don't write the comma for the last column
	for i := 1; i <= colCount; i++ {
		buffer.WriteString("?")
		if i != colCount {
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

func (this *sqlite3Storage) loadRow(tableName string, colCount int, values []string, db *sql.Tx, stmt *sql.Stmt, verbose bool) error {
	if len(values) == 0 || colCount == 0 {
		return nil
	}

	vals := make([]interface{}, 0)

	for i := 0; i < colCount; i++ {
		vals = append(vals, values[i])
	}

	_, err := stmt.Exec(vals...)

	if err != nil && verbose {
		fmt.Fprintln(os.Stderr, "Bad row: ", err)
	}

	return err
}

func (this *sqlite3Storage) ExecuteSQLStrings(sqlStrings []string) []*sql.Rows {
	var results []*sql.Rows

	for _, sqlQuery := range sqlStrings {
		if strings.Trim(sqlQuery, " ") != "" {
			result, err := this.db.Query(sqlQuery)
			if err != nil {
				log.Fatalln(err)
			}
			results = append(results, result)
		}
	}
	return results
}

func (this *sqlite3Storage) SaveTo(path string) {
	backup_db, open_err := sql.Open("sqlite3_textql", path)

	if open_err != nil {
		log.Fatalln(open_err)
	}

	backup_db.Ping()
	backupConnId := len(sqlite3conn) - 1

	// TODO: When backups are fixed in go-sqlite, remove debug logging and add error checking
	// https://github.com/mattn/go-sqlite3/issues/104

	log.Println("Start")

	backup, backup_start_err := sqlite3conn[backupConnId].Backup("main", sqlite3conn[this.connId], "main")

	if backup_start_err != nil {
		log.Fatalln(backup_start_err)
	}

	log.Println("Perform")

	backup_perform_error := backup.Step(-1)

	if backup_perform_error != nil {
		// 	log.Fatalln(backup_perform_error)
	}

	log.Println("Finish")

	backup_finish_error := backup.Finish()

	if backup_finish_error != nil {
		//log.Fatalln(backup_finish_error)
	}

	log.Println("Close")

	backup_close_error := backup_db.Close()

	if backup_close_error != nil {
		log.Fatalln(backup_close_error)
	}
}
