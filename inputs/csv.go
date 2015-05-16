package inputs

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type csvInput struct {
	options         *CSVInputOptions
	reader          *csv.Reader
	firstRow        []string
	header          []string
	minOutputLength int
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

	return this
}

func (this *csvInput) Name() string {
	if as_file, ok := this.options.ReadFrom.(*os.File); ok {
		if as_file.Name() == "/dev/stdin" {
			return "stdin"
		}
		return as_file.Name()
	}
	return "buffer"
}

func (this *csvInput) ReadRecord() []string {
	var row []string
	var file_err error

	if this.firstRow != nil {
		row = this.firstRow
		this.firstRow = nil
		return row
	}

	row, file_err = this.reader.Read()
	if file_err == io.EOF {
		return nil
	} else if parse_err, ok := file_err.(*csv.ParseError); ok {
		log.Println(parse_err)
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
	var read_err error

	this.firstRow, read_err = this.reader.Read()

	if read_err != nil {
		log.Fatalln(read_err)
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
