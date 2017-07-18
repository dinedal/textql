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
	whitespaceValuesCSV = `a,b,c
  , ,
1,2,3
4,5,6`
)

func NewTestCSVInput() (input inputs.Input, fp *os.File) {
	fp = test_util.OpenFileFromString(simpleCSV)

	opts := &inputs.CSVInputOptions{
		HasHeader: true,
		Separator: ',',
		ReadFrom:  fp,
	}

	newInput, _ := inputs.NewCSVInput(opts)
	return newInput, fp
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

	rows, rowsErr := storage.ExecuteSQLString(sqlString)

	if rowsErr != nil {
		t.Fatalf(rowsErr.Error())
	}

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

	rows, rowsErr := storage.ExecuteSQLString(sqlString)

	if rowsErr != nil {
		t.Fatalf(rowsErr.Error())
	}

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

func LoadTestDataAndExecuteQuery(t *testing.T, testData string, sqlString string) (map[int]map[string]interface{}, []string) {
	storage := NewSQLite3Storage(storageOpts)
	fp := test_util.OpenFileFromString(testData)

	opts := &inputs.CSVInputOptions{
		HasHeader: true,
		Separator: ',',
		ReadFrom:  fp,
	}

	input, _ := inputs.NewCSVInput(opts)
	defer fp.Close()
	defer os.Remove(fp.Name())
	defer storage.Close()

	storage.LoadInput(input)

	rows, rowsErr := storage.ExecuteSQLString(sqlString)

	if rowsErr != nil {
		t.Fatalf(rowsErr.Error())
	}

	cols, colsErr := rows.Columns()

	if colsErr != nil {
		t.Fatalf(colsErr.Error())
	}

	rowNumber := 0
	result := make(map[int]map[string]interface{})
	rawResult := make([]interface{}, len(cols))
	dest := make([]interface{}, len(cols))

	for i := range cols {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		scanErr := rows.Scan(dest...)

		if scanErr != nil {
			t.Fatalf(scanErr.Error())
		}

		result[rowNumber] = make(map[string]interface{})
		for i, raw := range rawResult {
			result[rowNumber][cols[i]] = raw
		}
		rowNumber++
	}

	return result, cols
}

func TestSQLiteStorageExecuteSQLStringMissingFromOuterQuery(t *testing.T) {
	data, cols := LoadTestDataAndExecuteQuery(t, simpleCSV, "count(*)")

	if len(cols) != 1 {
		t.Fatalf("Expected 1 column, got (%v)", len(cols))
	}

	intVal := data[0]["count(*)"].(int64)
	if intVal != 2 {
		t.Fatalf("Expected 2 rows counted, got (%v)", intVal)
	}
}

func TestSQLiteStorageExecuteSQLStringMissingFromSubQuery(t *testing.T) {
	data, cols := LoadTestDataAndExecuteQuery(t, simpleCSV, "count(*) from (select *)")

	if len(cols) != 1 {
		t.Fatalf("Expected 1 column, got (%v)", len(cols))
	}

	intVal := data[0]["count(*)"].(int64)
	if intVal != 2 {
		t.Fatalf("Expected 2 rows counted, got (%v)", intVal)
	}
}

func TestWhitespaceLoadsAsNull(t *testing.T) {
	data, cols := LoadTestDataAndExecuteQuery(t, whitespaceValuesCSV, "max(a)")

	if len(cols) != 1 {
		t.Fatalf("Expected 1 column, got (%v)", len(cols))
	}

	intVal := data[0]["max(a)"].(int64)
	if intVal != 4 {
		t.Fatalf("Expected 4 max value, got (%v)", intVal)
	}

	data, cols = LoadTestDataAndExecuteQuery(t, whitespaceValuesCSV, "typeof(a)")

	if len(cols) != 1 {
		t.Fatalf("Expected 1 column, got (%v)", len(cols))
	}

	uintVal := data[0]["typeof(a)"].([]uint8)
	if string(uintVal[:]) != "null" {
		t.Fatalf("Expected null value, got (%v)", uintVal)
	}

	uintVal = data[1]["typeof(a)"].([]uint8)
	if string(uintVal[:]) != "integer" {
		t.Fatalf("Expected integer value, got (%v)", uintVal)
	}

	uintVal = data[2]["typeof(a)"].([]uint8)
	if string(uintVal[:]) != "integer" {
		t.Fatalf("Expected integer value, got (%v)", uintVal)
	}

	data, cols = LoadTestDataAndExecuteQuery(t, whitespaceValuesCSV, "max(b)")

	if len(cols) != 1 {
		t.Fatalf("Expected 1 column, got (%v)", len(cols))
	}

	intVal = data[0]["max(b)"].(int64)
	if intVal != 5 {
		t.Fatalf("Expected 5 max value, got (%v)", intVal)
	}

	data, cols = LoadTestDataAndExecuteQuery(t, whitespaceValuesCSV, "typeof(b)")

	if len(cols) != 1 {
		t.Fatalf("Expected 1 column, got (%v)", len(cols))
	}

	uintVal = data[0]["typeof(b)"].([]uint8)
	if string(uintVal[:]) != "null" {
		t.Fatalf("Expected null value, got (%v)", uintVal)
	}

	uintVal = data[1]["typeof(b)"].([]uint8)
	if string(uintVal[:]) != "integer" {
		t.Fatalf("Expected integer value, got (%v)", uintVal)
	}

	uintVal = data[2]["typeof(b)"].([]uint8)
	if string(uintVal[:]) != "integer" {
		t.Fatalf("Expected integer value, got (%v)", uintVal)
	}
}
