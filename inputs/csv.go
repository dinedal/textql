package inputs

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// CSVInput represents a record producing input from a CSV formatted file or pipe.
type CSVInput struct {
	options         *CSVInputOptions
	reader          *csv.Reader
	firstRow        []string
	header          []string
	minOutputLength int
	name            string
}

// CSVInputOptions options are passed to the underlying encoding/csv reader.
type CSVInputOptions struct {
	// HasHeader when true, will treat the first row as a header row.
	HasHeader bool
	// Separator is the rune that fields are delimited by.
	Separator rune
	// ReadFrom is where the data will be read from.
	ReadFrom io.Reader
}

// NewCSVInput sets up a new CSVInput, the first row is read when this is run.
// If there is a problem with reading the first row, the error is returned.
// Otherwise, the returned csvInput can be reliably consumed with ReadRecord()
// until ReadRecord() returns nil.
func NewCSVInput(opts *CSVInputOptions) (*CSVInput, error) {
	csvInput := &CSVInput{
		options: opts,
		reader:  csv.NewReader(opts.ReadFrom),
	}
	csvInput.firstRow = nil

	csvInput.reader.FieldsPerRecord = -1
	csvInput.reader.Comma = csvInput.options.Separator
	csvInput.reader.LazyQuotes = true

	headerErr := csvInput.readHeader()

	if headerErr != nil {
		return nil, headerErr
	}

	if asFile, ok := csvInput.options.ReadFrom.(*os.File); ok {
		csvInput.name = filepath.Base(asFile.Name())
	} else {
		csvInput.name = "pipe"
	}

	return csvInput, nil
}

// Name returns the name of the CSV being read.
// By default, either the base filename or 'pipe' if it is a unix pipe
func (csvInput *CSVInput) Name() string {
	return csvInput.name
}

// SetName overrides the name of the CSV
func (csvInput *CSVInput) SetName(name string) {
	csvInput.name = name
}

// ReadRecord reads a single record from the CSV. Always returns successfully.
// If the record is empty, an empty []string is returned.
// Record expand to match the current row size, adding blank fields as needed.
// Records never return less then the number of fields in the first row.
// Returns nil on EOF
// In the event of a parse error due to an invalid record, it is logged, and
// an empty []string is returned with the number of fields in the first row,
// as if the record were empty.
//
// In general, this is a very tolerant of problems CSV reader.
func (csvInput *CSVInput) ReadRecord() []string {
	var row []string
	var fileErr error

	if csvInput.firstRow != nil {
		row = csvInput.firstRow
		csvInput.firstRow = nil
		return row
	}

	row, fileErr = csvInput.reader.Read()
	emptysToAppend := csvInput.minOutputLength - len(row)
	if fileErr == io.EOF {
		return nil
	} else if parseErr, ok := fileErr.(*csv.ParseError); ok {
		log.Println(parseErr)
		emptysToAppend = csvInput.minOutputLength
	}

	if emptysToAppend > 0 {
		for counter := 0; counter < emptysToAppend; counter++ {
			row = append(row, "")
		}
	}

	return row
}

func (csvInput *CSVInput) readHeader() error {
	var readErr error

	csvInput.firstRow, readErr = csvInput.reader.Read()

	if readErr != nil {
		log.Fatalln(readErr)
		return readErr
	}

	csvInput.minOutputLength = len(csvInput.firstRow)

	if csvInput.options.HasHeader {
		csvInput.header = csvInput.firstRow
		csvInput.firstRow = nil
	} else {
		csvInput.header = make([]string, csvInput.minOutputLength)
		for i := 0; i < len(csvInput.firstRow); i++ {
			csvInput.header[i] = "c" + strconv.Itoa(i)
		}
	}

	return nil
}

// Header returns the header of the csvInput. Either the first row if a header
// set in the options, or c#, where # is the column number, starting with 0.
func (csvInput *CSVInput) Header() []string {
	return csvInput.header
}
