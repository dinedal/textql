package storage

import (
	"bytes"
	"database/sql"
	"fmt"
	"path"
	"strings"

	"log"
	"os"
	"regexp"

	"github.com/dinedal/textql/inputs"
	"github.com/dinedal/textql/sqlparser"
	sqlite3 "github.com/mattn/go-sqlite3"
)

type sqlite3Storage struct {
	options        *SQLite3Options
	db             *sql.DB
	connId         int
	firstTableName string
}

type SQLite3Options struct{}

var (
	sqlite3conn []*sqlite3.SQLiteConn = []*sqlite3.SQLiteConn{}
)

func init() {
	sql.Register("sqlite3_textql",
		&sqlite3.SQLiteDriver{
			ConnectHook: func(conn *sqlite3.SQLiteConn) error {
				sqlite3conn = append(sqlite3conn, conn)
				return nil
			},
		})
}

func NewSQLite3Storage(opts *SQLite3Options) *sqlite3Storage {
	this := &sqlite3Storage{
		options:        opts,
		firstTableName: "",
	}

	this.open()
	return this
}

func (this *sqlite3Storage) open() {
	db, err := sql.Open("sqlite3_textql", ":memory:")

	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatalln(err)
	}

	this.connId = len(sqlite3conn) - 1
	this.db = db
}

func (this *sqlite3Storage) LoadInput(input inputs.Input) {
	tableName := strings.Replace(input.Name(), path.Ext(input.Name()), "", -1)
	this.createTable(tableName, input.Header(), false)

	tx, txErr := this.db.Begin()

	if txErr != nil {
		log.Fatalln(txErr)
	}

	stmt := this.createLoadStmt(tableName, len(input.Header()), tx)

	row := input.ReadRecord()
	for {
		if row == nil {
			break
		}
		this.loadRow(tableName, len(input.Header()), row, tx, stmt, true)
		row = input.ReadRecord()
	}
	stmt.Close()
	tx.Commit()

	if this.firstTableName == "" {
		this.firstTableName = tableName
	}
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

		buffer.WriteString("[" + col + "] NUMERIC")

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

	buffer.WriteString("INSERT INTO [" + (tableName) + "] VALUES (")
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

func (this *sqlite3Storage) ExecuteSQLString(sqlQuery string) *sql.Rows {
	var result *sql.Rows
	var err error

	if strings.Trim(sqlQuery, " ") != "" {
		implictFromSql := sqlparser.Magicify(sqlQuery, this.firstTableName)
		result, err = this.db.Query(implictFromSql)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return result
}

func (this *sqlite3Storage) SaveTo(path string) {
	backupDb, openErr := sql.Open("sqlite3_textql", path)

	if openErr != nil {
		log.Fatalln(openErr)
	}

	backupDb.Ping()
	backupConnId := len(sqlite3conn) - 1

	backup, backupStartErr := sqlite3conn[backupConnId].Backup("main", sqlite3conn[this.connId], "main")

	if backupStartErr != nil {
		log.Fatalln(backupStartErr)
	}

	_, backupPerformError := backup.Step(-1)

	if backupPerformError != nil {
		log.Fatalln(backupPerformError)
	}

	backupFinishError := backup.Finish()

	if backupFinishError != nil {
		log.Fatalln(backupFinishError)
	}

	backupCloseError := backupDb.Close()

	if backupCloseError != nil {
		log.Fatalln(backupCloseError)
	}
}

func (this *sqlite3Storage) Close() {
	this.db.Close()
}
