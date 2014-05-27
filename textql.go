package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/dinedal/textql/inputs"
	"github.com/dinedal/textql/storage"
	"github.com/dinedal/textql/util"
)

func main() {
	source_text := flag.String("source", "stdin", "Source file to load, or defaults to stdin")
	flag.Parse()

	fp := util.OpenFileOrStdin(source_text)

	opts := &inputs.CSVInputOptions{
		HasHeader: true,
		Seperator: ',',
		ReadFrom:  fp,
	}

	input := inputs.NewCSVInput(opts)

	fmt.Println(input.Name())

	storage_opts := &storage.SQLite3Options{}

	storage := storage.NewSQLite3Storage(storage_opts)

	storage.LoadInput(input)

	storage.ExecuteSQLStrings(strings.Split("select a from tbl;", ";"))

	storage.SaveTo("./out.db")

	// d := input.ReadRecord()
	// for {
	// 	if d == nil {
	// 		break
	// 	}
	// 	fmt.Println(d)
	// 	d = input.ReadRecord()
	// }

}
