package inputs

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

// SmartCSVInput represents a record producing input from a CSV formatted file or pipe.
// but it figures things out on it's own, if asked
type SmartCSVInput struct {
	options         *SmartCSVInputOptions
	reader          *csv.Reader
	rowBuffer       [][]string
	bufferLen       int
	bufferStart     int
	dataRemaining   bool
	headerTolerance int
	header          []string
	minOutputLength int
	name            string
}

// SmartCSVInputOptions options are passed to the underlying encoding/csv reader.
type SmartCSVInputOptions struct {
	// Seperator is the rune that fields are delimited by.
	Seperator rune
	// ReadFrom is where the data will be read from.
	ReadFrom io.Reader
	// Read-ahead buffer length
	MaxBufferLen int
}

func NewSmartCSVInput(opts *SmartCSVInputOptions) (*SmartCSVInput, error) {
	smartCSVInput := &SmartCSVInput{
		options: opts,
		reader:  csv.NewReader(opts.ReadFrom),
	}

	smartCSVInput.reader.FieldsPerRecord = -1
	smartCSVInput.reader.Comma = smartCSVInput.options.Seperator
	smartCSVInput.reader.LazyQuotes = true
	smartCSVInput.bufferStart = 0
	smartCSVInput.bufferLen = 0
	smartCSVInput.dataRemaining = true
	smartCSVInput.headerTolerance = 1

	if err := smartCSVInput.fillBuffer(); err != nil {
		return nil, err
	}

	if asFile, ok := smartCSVInput.options.ReadFrom.(*os.File); ok {
		smartCSVInput.name = path.Base(asFile.Name())
	} else {
		smartCSVInput.name = "pipe"
	}
	return smartCSVInput, nil
}

// SetHeaderTolerance allows users to adjust the header guess tolerance
func (smartCSVInput *SmartCSVInput) SetHeaderTolerance(tol int) {
	smartCSVInput.headerTolerance = tol
}

func (smartCSVInput *SmartCSVInput) fillBuffer() error {
	var row []string
	var fileErr error

	for smartCSVInput.bufferLen = 0; smartCSVInput.bufferLen <= smartCSVInput.options.MaxBufferLen; smartCSVInput.bufferLen++ {
		row, fileErr = smartCSVInput.reader.Read()
		if fileErr == io.EOF {
			smartCSVInput.dataRemaining = false
			return nil
		} else if parseErr, ok := fileErr.(*csv.ParseError); ok {
			log.Println(parseErr)
		}

		smartCSVInput.rowBuffer = append(smartCSVInput.rowBuffer, row)
	}

	return nil
}

func (smartCSVInput *SmartCSVInput) columnModalCount() int {
	counts := make(map[int]int)
	for _, row := range smartCSVInput.rowBuffer {
		length := 0
		for _, col := range row {
			if strings.TrimSpace(col) != "" {
				length++
			}
		}
		if length > 1 {
			counts[length]++
		}
	}

	maxModalCount := 0
	modalCount := 0
	for k, v := range counts {
		if v > maxModalCount {
			maxModalCount = v
			modalCount = k
		}
	}

	return modalCount
}

func (smartCSVInput *SmartCSVInput) headerGuess() []string {
	modalCount := smartCSVInput.columnModalCount()
	for i, row := range smartCSVInput.rowBuffer {
		length := 0
		for _, col := range row {
			if strings.TrimSpace(col) != "" {
				length++
			}
		}
		if length >= modalCount-smartCSVInput.headerTolerance {
			// TODO: type check for all strings in resultant row
			// also we move the buffer start past the non-header rows
			smartCSVInput.bufferStart = i
			smartCSVInput.minOutputLength = len(row)
			return row
		}
	}
	smartCSVInput.minOutputLength = len(smartCSVInput.rowBuffer[0])
	return smartCSVInput.rowBuffer[0]
}

// ReadRecord reads a single record from the CSV. Always returns successfully.
// If the record is empty, an empty []string is returned.
// Records expand to match the current row size, adding blank fields as needed.
// Records never return less then the number of fields in the header.
// Returns nil on EOF
// In the event of a parse error due to an invalid record, it is logged, and
// an empty []string is returned with the number of fields in the first row,
// as if the record were empty.
func (smartCSVInput *SmartCSVInput) ReadRecord() []string {
	var row []string
	var fileErr error

	row, fileErr = smartCSVInput.reader.Read()
	emptysToAppend := smartCSVInput.minOutputLength - len(row)
	if fileErr == io.EOF {
		return nil
	} else if parseErr, ok := fileErr.(*csv.ParseError); ok {
		log.Println(parseErr)
		emptysToAppend = smartCSVInput.minOutputLength
	}

	if emptysToAppend > 0 {
		for counter := 0; counter < emptysToAppend; counter++ {
			row = append(row, "")
		}
	}

	return row
}

// Header returns the contents of what SmartCSVInput considers to be the header row.
// If it is unable to guess a reasonable header, the first row is returned.
func (smartCSVInput *SmartCSVInput) Header() []string {
	if smartCSVInput.header == nil {
		smartCSVInput.header = smartCSVInput.headerGuess()
	}
	return smartCSVInput.header
}

// ClearHeader prevents the guessing of a header
func (smartCSVInput *SmartCSVInput) ClearHeader() {
	smartCSVInput.minOutputLength = 0
	smartCSVInput.header = make([]string, 0)
}

// Name returns the name of the CSV being read.
// By default, either the base filename or 'pipe' if it is a unix pipe
func (smartCSVInput *SmartCSVInput) Name() string {
	return smartCSVInput.name
}

// SetName overrides the name of the CSV
func (smartCSVInput *SmartCSVInput) SetName(name string) {
	smartCSVInput.name = name
}
