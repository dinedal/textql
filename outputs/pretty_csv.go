package outputs

import (
	"database/sql"
	"io"
	"log"

	"github.com/olekukonko/tablewriter"
)

// PrettyCSVOutput represents a TextQL output that transforms sql.Rows into pretty tables
type PrettyCSVOutput struct {
	options         *PrettyCSVOutputOptions
	writer          *tablewriter.Table
	firstRow        []string
	header          []string
	minOutputLength int
}

// PrettyCSVOutputOptions define options that are passed to tablewriter for formatting
// the output in specific ways.
type PrettyCSVOutputOptions struct {
	// WriteHeader determines if a header row based on the column names should be written.
	WriteHeader bool
	// WriteTo is where the formatted data will be written to.
	WriteTo io.Writer
}

// NewPrettyCSVOutput returns a new PrettyCSVOutput configured per the options provided.
func NewPrettyCSVOutput(opts *PrettyCSVOutputOptions) *PrettyCSVOutput {
	prettyCsvOutput := &PrettyCSVOutput{
		options: opts,
		writer:  tablewriter.NewWriter(opts.WriteTo),
	}

	return prettyCsvOutput
}

// Show writes the sql.Rows given to the destination in tablewriter basic format.
func (prettyCsvOutput *PrettyCSVOutput) Show(rows *sql.Rows) {
	cols, colsErr := rows.Columns()

	if colsErr != nil {
		log.Fatalln(colsErr)
	}

	if prettyCsvOutput.options.WriteHeader {
		prettyCsvOutput.writer.SetHeader(cols)
		prettyCsvOutput.writer.SetHeaderLine(true)
		prettyCsvOutput.writer.SetAutoFormatHeaders(false)
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

		prettyCsvOutput.writer.Append(result)
	}

	if len(cols) > 0 {
		prettyCsvOutput.writer.Render()
	}
	rows.Close()
}
