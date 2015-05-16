package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/dinedal/textql/inputs"
	"github.com/dinedal/textql/outputs"
	"github.com/dinedal/textql/storage"
	"github.com/dinedal/textql/util"
)

func main() {
	commands := flag.String("sql", "", "SQL Command(s) to run on the data")
	sourceFile := flag.String("source", "stdin", "Source file to load, or defaults to stdin")
	delimiter := flag.String("dlm", ",", "Input delimiter between fields -dlm=tab for tab, -dlm=0x## to specify a character code in hex")
	header := flag.Bool("header", false, "Treat file as having the first row as a header row")
	outputHeader := flag.Bool("output-header", false, "Display column names in output")
	outputDelimiter := flag.String("output-dlm", ",", "Output delimiter between fields -output-dlm=tab for tab, -dlm=0x## to specify a character code in hex")
	outputFile := flag.String("output-file", "stdout", "Filename to write output to, if empty no output is written")
	tableName := flag.String("table-name", "", "Override the default table name (input file name or stdin)")
	saveTo := flag.String("save-to", "", "If set, sqlite3 db is left on disk at this path")
	console := flag.Bool("console", false, "After all commands are run, open sqlite3 console with this data")
	flag.Parse()

	if *console {
		if *sourceFile == "stdin" {
			log.Fatalln("Can not open console with pipe input, read a file instead")
		}
		_, sqlite3ConsolePathErr := exec.LookPath("sqlite3")
		if sqlite3ConsolePathErr != nil {
			log.Fatalln("Console requested but unable to locate `sqlite3` program on $PATH")
		}
	}

	fp := util.OpenFileOrStdDev(sourceFile)

	opts := &inputs.CSVInputOptions{
		HasHeader: *header,
		Seperator: util.DetermineSeparator(delimiter),
		ReadFrom:  fp,
	}

	input := inputs.NewCSVInput(opts)

	storage_opts := &storage.SQLite3Options{}

	storage := storage.NewSQLite3Storage(storage_opts)

	if (*tableName) != "" {
		storage.LoadInput(input, *tableName)
	} else {
		storage.LoadInput(input, path.Base(input.Name()))
	}

	queryResults := storage.ExecuteSQLStrings(strings.Split(*commands, ";"))

	if (*outputFile) != "" {
		displayOpts := &outputs.CSVOutputOptions{
			WriteHeader: *outputHeader,
			Seperator:   util.DetermineSeparator(outputDelimiter),
			WriteTo:     util.OpenFileOrStdDev(outputFile),
		}

		outputer := outputs.NewCSVOutput(displayOpts)
		outputer.Show(queryResults)
	}

	if (*saveTo) != "" {
		storage.SaveTo(*util.CleanPath(saveTo))
	}

	if *console {
		var args []string

		if *outputHeader {
			args = []string{"-header"}
		} else {
			args = []string{}
		}

		if (*saveTo) != "" {
			args = append(args, *util.CleanPath(saveTo))
		} else {
			tempFile, err := ioutil.TempFile(os.TempDir(), "textql")
			if err != nil {
				log.Fatalln(err)
			}
			defer os.Remove(tempFile.Name())
			tempFile.Close()
			storage.SaveTo(tempFile.Name())
			args = append(args, tempFile.Name())
		}

		cmd := exec.Command("sqlite3", args...)

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd_err := cmd.Run()

		log.Println("ok whut")

		if cmd.Process != nil {
			cmd.Process.Release()
		}

		if cmd_err != nil {
			log.Fatalln(cmd_err)
		}
	} else {
		storage.Close()
	}
}
