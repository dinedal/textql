package main

import (
	"flag"
	"fmt"

	"github.com/dinedal/textql/inputs"
	"github.com/dinedal/textql/util"
)

func main() {
	source_text := flag.String("source", "stdin", "Source file to load, or defaults to stdin")
	flag.Parse()

	fp := util.OpenFileOrStdin(source_text)

	opts := &inputs.CSVInputOptions{
		HasHeader: false,
		Seperator: ',',
		ReadFrom:  fp,
	}

	input := inputs.NewCSVInput(opts)

	fmt.Println(input.Name())

	d := input.ReadRecord()
	for {
		if d == nil {
			break
		}
		fmt.Println(d)
		d = input.ReadRecord()
	}
}
