package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/dinedal/textql/inputs"
	"github.com/dinedal/textql/outputs"
	"github.com/dinedal/textql/storage"
	"github.com/dinedal/textql/util"
)

type CommandLineOptions struct {
	Commands        *string
	SourceFile      *string
	Delimiter       *string
	Header          *bool
	OutputHeader    *bool
	OutputDelimiter *string
	OutputFile      *string
	SaveTo          *string
	Console         *bool
}

func NewCommandLineOptions() *CommandLineOptions {
	cmdLineOpts := CommandLineOptions{}
	cmdLineOpts.Commands = flag.String("sql", "", "SQL Command(s) to run on the data")
	cmdLineOpts.SourceFile = flag.String("source", "stdin", "Source file or directory to load, or defaults to stdin")
	cmdLineOpts.Delimiter = flag.String("dlm", ",", "Input delimiter between fields -dlm=tab for tab, -dlm=0x## to specify a character code in hex")
	cmdLineOpts.Header = flag.Bool("header", false, "Treat file as having the first row as a header row")
	cmdLineOpts.OutputHeader = flag.Bool("output-header", false, "Display column names in output")
	cmdLineOpts.OutputDelimiter = flag.String("output-dlm", ",", "Output delimiter between fields -output-dlm=tab for tab, -dlm=0x## to specify a character code in hex")
	cmdLineOpts.OutputFile = flag.String("output-file", "stdout", "Filename to write output to, if empty no output is written")
	cmdLineOpts.SaveTo = flag.String("save-to", "", "If set, sqlite3 db is left on disk at this path")
	cmdLineOpts.Console = flag.Bool("console", false, "After all commands are run, open sqlite3 console with this data")
	flag.Usage = cmdLineOpts.Usage
	flag.Parse()

	return &cmdLineOpts
}

func (this *CommandLineOptions) GetCommands() string {
	return *this.Commands
}

func (this *CommandLineOptions) GetSourceFile() string {
	return *this.SourceFile
}

func (this *CommandLineOptions) GetSourceFiles() []string {
	return append(flag.Args(), *this.SourceFile)
}

func (this *CommandLineOptions) GetDelimiter() string {
	return *this.Delimiter
}

func (this *CommandLineOptions) GetHeader() bool {
	return *this.Header
}

func (this *CommandLineOptions) GetOutputHeader() bool {
	return *this.OutputHeader
}

func (this *CommandLineOptions) GetOutputDelimiter() string {
	return *this.OutputDelimiter
}

func (this *CommandLineOptions) GetOutputFile() string {
	return *this.OutputFile
}

func (this *CommandLineOptions) GetSaveTo() string {
	return util.CleanPath(*this.SaveTo)
}

func (this *CommandLineOptions) GetConsole() bool {
	return *this.Console
}

func (this *CommandLineOptions) Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "  %s [-console] [-save-to path] [-output-file path] [-output-dlm] [-output-header] [-header] [-dlm delimter] [-source path] [-sql sql_statements] [path ...] \n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\n")
	flag.PrintDefaults()
}

func main() {
	cmdLineOpts := NewCommandLineOptions()
	var outputer outputs.Output

	if cmdLineOpts.GetConsole() {
		if cmdLineOpts.GetSourceFile() == "stdin" && util.IsThereDataOnStdin() {
			log.Fatalln("Can not open console with pipe input, read a file instead")
		}
		_, sqlite3ConsolePathErr := exec.LookPath("sqlite3")
		if sqlite3ConsolePathErr != nil {
			log.Fatalln("Console requested but unable to locate `sqlite3` program on $PATH")
		}
	}

	inputSources := make([]string, 0)

	for _, sourceFile := range cmdLineOpts.GetSourceFiles() {
		if sourceFile == "stdin" && util.IsThereDataOnStdin() {
			inputSources = append(inputSources, sourceFile)
		} else if sourceFile != "stdin" {
			if util.IsPathDir(sourceFile) {
				for _, file := range util.AllFilesInDirectory(sourceFile) {
					inputSources = append(inputSources, file)
				}
			} else {
				inputSources = append(inputSources, sourceFile)
			}
		}
	}

	storageOpts := &storage.SQLite3Options{}

	storage := storage.NewSQLite3Storage(storageOpts)

	for _, file := range inputSources {
		fp := util.OpenFileOrStdDev(file)

		inputOpts := &inputs.CSVInputOptions{
			HasHeader: cmdLineOpts.GetHeader(),
			Seperator: util.DetermineSeparator(cmdLineOpts.GetDelimiter()),
			ReadFrom:  fp,
		}

		input := inputs.NewCSVInput(inputOpts)

		storage.LoadInput(input)
	}

	sqlStrings := strings.Split(cmdLineOpts.GetCommands(), ";")

	if cmdLineOpts.GetOutputFile() != "" {
		displayOpts := &outputs.CSVOutputOptions{
			WriteHeader: cmdLineOpts.GetOutputHeader(),
			Seperator:   util.DetermineSeparator(cmdLineOpts.GetOutputDelimiter()),
			WriteTo:     util.OpenFileOrStdDev(cmdLineOpts.GetOutputFile()),
		}

		outputer = outputs.NewCSVOutput(displayOpts)
	}

	for _, sqlQuery := range sqlStrings {
		queryResults := storage.ExecuteSQLString(sqlQuery)

		if queryResults != nil && cmdLineOpts.GetOutputFile() != "" {
			outputer.Show(queryResults)
		}
	}

	if cmdLineOpts.GetSaveTo() != "" {
		storage.SaveTo(cmdLineOpts.GetSaveTo())
	}

	if cmdLineOpts.GetConsole() {
		var args []string

		if cmdLineOpts.GetOutputHeader() {
			args = []string{"-header"}
		} else {
			args = []string{}
		}

		if cmdLineOpts.GetSaveTo() != "" {
			args = append(args, cmdLineOpts.GetSaveTo())
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
		cmdErr := cmd.Run()

		if cmd.Process != nil {
			cmd.Process.Release()
		}

		if cmdErr != nil {
			log.Fatalln(cmdErr)
		}
	} else {
		storage.Close()
	}
}
