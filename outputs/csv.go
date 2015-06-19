package outputs

import (
	"database/sql"
	"encoding/csv"
	"io"
	"log"
)

type csvOutput struct {
	options         *CSVOutputOptions
	writer          *csv.Writer
	firstRow        []string
	header          []string
	minOutputLength int
}

type CSVOutputOptions struct {
	WriteHeader bool
	Seperator   rune
	WriteTo     io.Writer
}

func NewCSVOutput(opts *CSVOutputOptions) *csvOutput {
	this := &csvOutput{
		options: opts,
		writer:  csv.NewWriter(opts.WriteTo),
	}

	this.writer.Comma = this.options.Seperator

	return this
}

func (this *csvOutput) Show(rows *sql.Rows) {
	cols, colsErr := rows.Columns()

	if colsErr != nil {
		log.Fatalln(colsErr)
	}

	if this.options.WriteHeader {
		if err := this.writer.Write(cols); err != nil {
			log.Fatalln(err)
		}
	}

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

		writeErr := this.writer.Write(result)

		if writeErr != nil {
			log.Fatalln(colsErr)
		}
	}

	this.writer.Flush()
	rows.Close()
}
