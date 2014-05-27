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

	fmt.Println(fp.Name())

	opts := &inputs.CSVInputOptions{
		HasHeader: false,
		Seperator: ',',
		ReadFrom:  fp,
	}

	input := inputs.NewCSVInput(opts)

	d := input.ReadRecord()
	for {
		if d == nil {
			break
		}
		fmt.Println(d)
		d = input.ReadRecord()
	}
}
