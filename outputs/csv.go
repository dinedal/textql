package outputs

import (
	"database/sql"
	"encoding/csv"
	"io"
	"log"
)

// CSVOutput represents a TextQL output that transforms sql.Rows into CSV formatted
// string data using encoding/csv
type CSVOutput struct {
	options         *CSVOutputOptions
	writer          *csv.Writer
	firstRow        []string
	header          []string
	minOutputLength int
}

// CSVOutputOptions define options that are passed to encoding/csv for formatting
// the output in specific ways.
type CSVOutputOptions struct {
	// WriteHeader determines if a header row based on the column names should be written.
	WriteHeader bool
	// Separator is the rune used to delimit fields.
	Separator rune
	// WriteTo is where the formatted data will be written to.
	WriteTo io.Writer
}

// NewCSVOutput returns a new CSVOutput configured per the options provided.
func NewCSVOutput(opts *CSVOutputOptions) *CSVOutput {
	csvOutput := &CSVOutput{
		options: opts,
		writer:  csv.NewWriter(opts.WriteTo),
	}

	csvOutput.writer.Comma = csvOutput.options.Separator

	return csvOutput
}

// Show writes the sql.Rows given to the destination in CSV format.
func (csvOutput *CSVOutput) Show(rows *sql.Rows) {
	cols, colsErr := rows.Columns()

	if colsErr != nil {
		log.Fatalln(colsErr)
	}

	if csvOutput.options.WriteHeader && len(cols) > 0 {
		if err := csvOutput.writer.Write(cols); err != nil {
			log.Fatalln(err)
		}
	}

	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))

	dest := make([]interface{}, len(cols))

	for i := range cols {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		rows.Scan(dest...)

		for i, raw := range rawResult {
			result[i] = string(raw)
		}

		writeErr := csvOutput.writer.Write(result)

		if writeErr != nil {
			log.Fatalln(writeErr)
		}
	}

	csvOutput.writer.Flush()
	rows.Close()
}
