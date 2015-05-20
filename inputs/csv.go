package inputs

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path"
	"strconv"
)

type csvInput struct {
	options         *CSVInputOptions
	reader          *csv.Reader
	firstRow        []string
	header          []string
	minOutputLength int
	name            string
}

type CSVInputOptions struct {
	HasHeader bool
	Seperator rune
	ReadFrom  io.Reader
}

func NewCSVInput(opts *CSVInputOptions) *csvInput {
	this := &csvInput{
		options: opts,
		reader:  csv.NewReader(opts.ReadFrom),
	}
	this.firstRow = nil

	this.reader.FieldsPerRecord = -1
	this.reader.Comma = this.options.Seperator
	this.reader.LazyQuotes = true

	this.readHeader()

	if asFile, ok := this.options.ReadFrom.(*os.File); ok {
		this.name = path.Base(asFile.Name())
	} else {
		this.name = "pipe"
	}

	return this
}

func (this *csvInput) Name() string {
	return this.name
}

func (this *csvInput) SetName(name string) {
	this.name = name
}

func (this *csvInput) ReadRecord() []string {
	var row []string
	var fileErr error

	if this.firstRow != nil {
		row = this.firstRow
		this.firstRow = nil
		return row
	}

	row, fileErr = this.reader.Read()
	if fileErr == io.EOF {
		return nil
	} else if parseErr, ok := fileErr.(*csv.ParseError); ok {
		log.Println(parseErr)
	}
	emptysToAppend := this.minOutputLength - len(row)

	if emptysToAppend > 0 {
		for counter := 0; counter < emptysToAppend; counter++ {
			row = append(row, "")
		}
	}

	return row
}

func (this *csvInput) readHeader() {
	var readErr error

	this.firstRow, readErr = this.reader.Read()

	if readErr != nil {
		log.Fatalln(readErr)
	}

	this.minOutputLength = len(this.firstRow)

	if this.options.HasHeader {
		this.header = this.firstRow
		this.firstRow = nil
	} else {
		this.header = make([]string, this.minOutputLength)
		for i := 0; i < len(this.firstRow); i++ {
			this.header[i] = "c" + strconv.Itoa(i)
		}
	}
}

func (this *csvInput) Header() []string {
	return this.header
}
