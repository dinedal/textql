package storage

import (
	"bytes"
	"database/sql"
	"path"
	"strings"

	"log"
	"regexp"

	"github.com/dinedal/textql/inputs"
	"github.com/dinedal/textql/sqlparser"

	sqlite3 "github.com/mattn/go-sqlite3"
)

// SQLite3Storage represents a TextQL compatible SQL backend based on in-memory SQLite3
type SQLite3Storage struct {
	options        *SQLite3Options
	db             *sql.DB
	connID         int
	firstTableName string
}

// SQLite3Options are options passed into SQLite3 connection as needed.
type SQLite3Options struct{}

var (
	sqlite3conn          = []*sqlite3.SQLiteConn{}
	allWhiteSpace        = regexp.MustCompile("^\\s+$")
	tableNameCheckRegEx  = regexp.MustCompile(`.*\[.*\].*`)
	columnNameCheckRegEx = regexp.MustCompile(`.*\[.*\].*`)
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

// NewSQLite3StorageWithDefaults returns a SQLite3Storage with the default options.
func NewSQLite3StorageWithDefaults() *SQLite3Storage {
	return NewSQLite3Storage(&SQLite3Options{})
}

// NewSQLite3Storage returns a SQLite3Storage with the SQLite3Options provided applied.
func NewSQLite3Storage(opts *SQLite3Options) *SQLite3Storage {
	sqlite3Storage := &SQLite3Storage{
		options:        opts,
		firstTableName: "",
	}

	sqlite3Storage.open()
	return sqlite3Storage
}

func (sqlite3Storage *SQLite3Storage) open() {
	db, err := sql.Open("sqlite3_textql", ":memory:")

	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatalln(err)
	}

	sqlite3Storage.connID = len(sqlite3conn) - 1
	sqlite3Storage.db = db
}

// LoadInput reads the entire Input provided into a table named after the Input name.
// The name is cooreced into a valid SQLite3 table name prior to use.
func (sqlite3Storage *SQLite3Storage) LoadInput(input inputs.Input) {
	tableName := strings.Replace(input.Name(), path.Ext(input.Name()), "", -1)
	sqlite3Storage.createTable(tableName, input.Header(), false)

	tx, txErr := sqlite3Storage.db.Begin()

	if txErr != nil {
		log.Fatalln(txErr)
	}

	stmt := sqlite3Storage.createLoadStmt(tableName, len(input.Header()), tx)

	row := input.ReadRecord()
	for {
		if row == nil {
			break
		}
		sqlite3Storage.loadRow(tableName, len(input.Header()), row, tx, stmt, true)
		row = input.ReadRecord()
	}
	stmt.Close()
	tx.Commit()

	if sqlite3Storage.firstTableName == "" {
		sqlite3Storage.firstTableName = tableName
	}
}

func (sqlite3Storage *SQLite3Storage) createTable(tableName string, columnNames []string, verbose bool) error {
	var buffer bytes.Buffer

	if tableNameCheckRegEx.FindString(tableName) != "" {
		log.Fatalln("Invalid table name", tableName)
	}

	buffer.WriteString("CREATE TABLE IF NOT EXISTS [" + (tableName) + "] (")

	for i, col := range columnNames {
		if columnNameCheckRegEx.FindString(col) != "" {
			log.Fatalln("Invalid table name", col)
		}

		buffer.WriteString("[" + col + "] NUMERIC")

		if i != len(columnNames)-1 {
			buffer.WriteString(", ")
		}
	}

	buffer.WriteString(");")

	_, err := sqlite3Storage.db.Exec(buffer.String())

	if err != nil {
		log.Fatalln(err)
	}

	if verbose {
		log.Println(buffer.String())
	}

	return err
}

func (sqlite3Storage *SQLite3Storage) createLoadStmt(tableName string, colCount int, db *sql.Tx) *sql.Stmt {
	if colCount == 0 {
		log.Fatalln("Nothing to build insert with!")
	}
	var buffer bytes.Buffer

	buffer.WriteString("INSERT INTO [" + (tableName) + "] VALUES (")
	// Don't write the comma for the last column
	for i := 1; i <= colCount; i++ {
		buffer.WriteString("nullif(?,'')")
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

func (sqlite3Storage *SQLite3Storage) loadRow(tableName string, colCount int, values []string, db *sql.Tx, stmt *sql.Stmt, verbose bool) error {
	if len(values) == 0 || colCount == 0 {
		return nil
	}

	var vals []interface{}

	for i := 0; i < colCount; i++ {
		if allWhiteSpace.MatchString(values[i]) {
			vals = append(vals, "")
		} else {
			vals = append(vals, values[i])
		}
	}

	_, err := stmt.Exec(vals...)

	if err != nil && verbose {
		log.Printf("Bad row: %v\n", err)
	}

	return err
}

// ExecuteSQLString maps the sqlQuery provided from short hand TextQL to SQL, then
// applies the query to the sqlite3 in memory database, and lastly returns the sql.Rows
// that resulted from the executing query.
func (sqlite3Storage *SQLite3Storage) ExecuteSQLString(sqlQuery string) (*sql.Rows, error) {
	var result *sql.Rows
	var err error

	if strings.Trim(sqlQuery, " ") != "" {
		implictFromSQL := sqlparser.Magicify(sqlQuery, sqlite3Storage.firstTableName)
		result, err = sqlite3Storage.db.Query(implictFromSQL)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// SaveTo saves the current in memory database to the path provided as a string.
func (sqlite3Storage *SQLite3Storage) SaveTo(path string) error {
	backupDb, openErr := sql.Open("sqlite3_textql", path)
	if openErr != nil {
		return openErr
	}

	backupPingErr := backupDb.Ping()
	if backupPingErr != nil {
		return backupPingErr
	}
	backupConnID := len(sqlite3conn) - 1

	backup, backupStartErr := sqlite3conn[backupConnID].Backup("main", sqlite3conn[sqlite3Storage.connID], "main")
	if backupStartErr != nil {
		return backupStartErr
	}

	_, backupPerformError := backup.Step(-1)
	if backupPerformError != nil {
		return backupPerformError
	}

	backupFinishError := backup.Finish()
	if backupFinishError != nil {
		return backupFinishError
	}

	backupCloseError := backupDb.Close()
	if backupCloseError != nil {
		return backupCloseError
	}

	return nil
}

// Close will close the current database
func (sqlite3Storage *SQLite3Storage) Close() {
	sqlite3Storage.db.Close()
}
