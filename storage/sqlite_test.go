package storage

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/dinedal/textql/inputs"
	"github.com/dinedal/textql/test_util"
)

var (
	storageOpts = &SQLite3Options{}
	simpleCSV   = `a,b,c
1,2,3
4,5,6`
)

func NewTestCSVInput() (input inputs.Input, fp *os.File) {
	fp = test_util.OpenFileFromString(simpleCSV)

	opts := &inputs.CSVInputOptions{
		HasHeader: true,
		Seperator: ',',
		ReadFrom:  fp,
	}

	return inputs.NewCSVInput(opts), fp
}

func TestSQLiteStorageLoadInput(t *testing.T) {
	storage := NewSQLite3Storage(storageOpts)
	input, fp := NewTestCSVInput()
	defer fp.Close()
	defer os.Remove(fp.Name())
	defer storage.Close()

	storage.LoadInput(input)
}

func TestSQLiteStorageSaveTo(t *testing.T) {
	var (
		cmdOut   []byte
		err      error
		tempFile *os.File
	)

	storage := NewSQLite3Storage(storageOpts)
	input, fp := NewTestCSVInput()
	defer fp.Close()
	defer os.Remove(fp.Name())
	defer storage.Close()

	storage.LoadInput(input)

	tempFile, err = ioutil.TempFile(os.TempDir(), "textql_test")

	if err != nil {
		t.Fatalf(err.Error())
	}

	defer os.Remove(tempFile.Name())
	tempFile.Close()
	storage.SaveTo(tempFile.Name())
	storage.Close()

	args := []string{tempFile.Name(), "pragma integrity_check;"}

	cmd := exec.Command("sqlite3", args...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if cmdOut, err = cmd.Output(); err != nil {
		fmt.Println(string(cmdOut))
		fmt.Println(args)
		t.Fatalf(err.Error())
	}
	cmdOutStr := string(cmdOut)

	if cmdOutStr != "ok\n" {
		fmt.Println(cmdOutStr)
		t.Fatalf("SaveTo integrity check failed!")
	}
}

func TestSQLiteStorageExecuteSQLStringNormalSQL(t *testing.T) {
	storage := NewSQLite3Storage(storageOpts)
	input, fp := NewTestCSVInput()
	defer fp.Close()
	defer os.Remove(fp.Name())
	defer storage.Close()

	storage.LoadInput(input)

	sqlString := "select count(*) from " + storage.firstTableName

	rows := storage.ExecuteSQLString(sqlString)

	cols, colsErr := rows.Columns()

	if colsErr != nil {
		t.Fatalf(colsErr.Error())
	}

	if len(cols) != 1 {
		t.Fatalf("Expected 1 column, got (%v)", len(cols))
	}

	var dest int

	for rows.Next() {
		rows.Scan(&dest)
		if dest != 2 {
			t.Fatalf("Expected 2 rows counted, got (%v)", dest)
		}
	}
}

func TestSQLiteStorageExecuteSQLStringMissingSelect(t *testing.T) {
	storage := NewSQLite3Storage(storageOpts)
	input, fp := NewTestCSVInput()
	defer fp.Close()
	defer os.Remove(fp.Name())
	defer storage.Close()

	storage.LoadInput(input)

	sqlString := "count(*) from " + storage.firstTableName

	rows := storage.ExecuteSQLString(sqlString)

	cols, colsErr := rows.Columns()

	if colsErr != nil {
		t.Fatalf(colsErr.Error())
	}

	if len(cols) != 1 {
		t.Fatalf("Expected 1 column, got (%v)", len(cols))
	}

	var dest int

	for rows.Next() {
		rows.Scan(&dest)
		if dest != 2 {
			t.Fatalf("Expected 2 rows counted, got (%v)", dest)
		}
	}
}

func TestSQLiteStorageExecuteSQLStringMissingFromOuterQuery(t *testing.T) {
	storage := NewSQLite3Storage(storageOpts)
	input, fp := NewTestCSVInput()
	defer fp.Close()
	defer os.Remove(fp.Name())
	defer storage.Close()

	storage.LoadInput(input)

	sqlString := "count(*)"

	rows := storage.ExecuteSQLString(sqlString)

	cols, colsErr := rows.Columns()

	if colsErr != nil {
		t.Fatalf(colsErr.Error())
	}

	if len(cols) != 1 {
		t.Fatalf("Expected 1 column, got (%v)", len(cols))
	}

	var dest int

	for rows.Next() {
		rows.Scan(&dest)
		if dest != 2 {
			t.Fatalf("Expected 2 rows counted, got (%v)", dest)
		}
	}
}

func TestSQLiteStorageExecuteSQLStringMissingFromSubQuery(t *testing.T) {
	storage := NewSQLite3Storage(storageOpts)
	input, fp := NewTestCSVInput()
	defer fp.Close()
	defer os.Remove(fp.Name())
	defer storage.Close()

	storage.LoadInput(input)

	sqlString := "count(*) from (select *)"

	rows := storage.ExecuteSQLString(sqlString)

	cols, colsErr := rows.Columns()

	if colsErr != nil {
		t.Fatalf(colsErr.Error())
	}

	if len(cols) != 1 {
		t.Fatalf("Expected 1 column, got (%v)", len(cols))
	}

	var dest int

	for rows.Next() {
		rows.Scan(&dest)
		if dest != 2 {
			t.Fatalf("Expected 2 rows counted, got (%v)", dest)
		}
	}
}
