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

type commandLineOptions struct {
	Statements      *string
	SourceFile      *string
	Delimiter       *string
	Header          *bool
	OutputHeader    *bool
	OutputDelimiter *string
	OutputFile      *string
	SaveTo          *string
	Console         *bool
	Version         *bool
	Quiet           *bool
	Pretty          *bool
}

// Must be set at build via -ldflags "-X main.VERSION=`cat VERSION`"
var VERSION string

func newCommandLineOptions() *commandLineOptions {
	cmdLineOpts := commandLineOptions{}
	cmdLineOpts.Statements = flag.String("sql", "", "SQL Statement(s) to run on the data")
	cmdLineOpts.Delimiter = flag.String("dlm", ",", "Input delimiter character between fields -dlm=tab for tab, -dlm=0x## to specify a character code in hex")
	cmdLineOpts.Header = flag.Bool("header", false, "Treat input files as having the first row as a header row")
	cmdLineOpts.OutputHeader = flag.Bool("output-header", false, "Display column names in output")
	cmdLineOpts.OutputDelimiter = flag.String("output-dlm", ",", "Output delimiter character between fields -output-dlm=tab for tab, -dlm=0x## to specify a character code in hex")
	cmdLineOpts.OutputFile = flag.String("output-file", "stdout", "Filename to write output to, if empty no output is written")
	cmdLineOpts.SaveTo = flag.String("save-to", "", "SQLite3 db is left on disk at this file")
	cmdLineOpts.Console = flag.Bool("console", false, "After all statements are run, open SQLite3 REPL with this data")
	cmdLineOpts.Version = flag.Bool("version", false, "Print version and exit")
	cmdLineOpts.Quiet = flag.Bool("quiet", false, "Surpress logging")
	cmdLineOpts.Pretty = flag.Bool("pretty", false, "Output pretty formatting")
	flag.Usage = cmdLineOpts.Usage
	flag.Parse()

	return &cmdLineOpts
}

func (clo *commandLineOptions) GetStatements() string {
	return *clo.Statements
}

func (clo *commandLineOptions) GetSourceFiles() []string {
	return flag.Args()
}

func (clo *commandLineOptions) GetDelimiter() string {
	return *clo.Delimiter
}

func (clo *commandLineOptions) GetHeader() bool {
	return *clo.Header
}

func (clo *commandLineOptions) GetOutputHeader() bool {
	return *clo.OutputHeader
}

func (clo *commandLineOptions) GetOutputDelimiter() string {
	return *clo.OutputDelimiter
}

func (clo *commandLineOptions) GetOutputFile() string {
	return *clo.OutputFile
}

func (clo *commandLineOptions) GetSaveTo() string {
	return util.CleanPath(*clo.SaveTo)
}

func (clo *commandLineOptions) GetConsole() bool {
	return *clo.Console
}

func (clo *commandLineOptions) GetVersion() bool {
	return *clo.Version
}

func (clo *commandLineOptions) GetQuiet() bool {
	return *clo.Quiet
}

func (clo *commandLineOptions) GetPretty() bool {
	return *clo.Pretty
}

func (clo *commandLineOptions) Usage() {
	if !clo.GetQuiet() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "  %s [-console] [-save-to path path] [-output-file path] [-output-dlm delimter] [-output-header] [-pretty] [-quiet] [-header] [-dlm delimter] [-sql sql_statements] [path ...] \n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\n")
		flag.PrintDefaults()
	}
}

func main() {
	cmdLineOpts := newCommandLineOptions()
	var outputer outputs.Output

	if cmdLineOpts.GetVersion() {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	if len(cmdLineOpts.GetSourceFiles()) == 0 && !util.IsThereDataOnStdin() {
		cmdLineOpts.Usage()
	}

	if cmdLineOpts.GetQuiet() {
		log.SetOutput(ioutil.Discard)
	}

	if cmdLineOpts.GetConsole() {
		if util.IsThereDataOnStdin() {
			log.Fatalln("Can not open console with pipe input, read a file instead")
		}
		_, sqlite3ConsolePathErr := exec.LookPath("sqlite3")
		if sqlite3ConsolePathErr != nil {
			log.Fatalln("Console requested but unable to locate `sqlite3` program on $PATH")
		}
	}

	var inputSources []string

	if util.IsThereDataOnStdin() {
		inputSources = append(inputSources, "stdin")
	}

	for _, sourceFile := range cmdLineOpts.GetSourceFiles() {
		if util.IsPathDir(sourceFile) {
			for _, file := range util.AllFilesInDirectory(sourceFile) {
				inputSources = append(inputSources, file)
			}
		} else {
			inputSources = append(inputSources, sourceFile)
		}
	}

	storage := storage.NewSQLite3StorageWithDefaults()

	for _, file := range inputSources {
		fp := util.OpenFileOrStdDev(file, false)

		inputOpts := &inputs.CSVInputOptions{
			HasHeader: cmdLineOpts.GetHeader(),
			Seperator: util.DetermineSeparator(cmdLineOpts.GetDelimiter()),
			ReadFrom:  fp,
		}

		input, inputErr := inputs.NewCSVInput(inputOpts)

		if inputErr != nil {
			log.Printf("Unable to load %v\n", file)
		}

		storage.LoadInput(input)
	}

	sqlStrings := strings.Split(cmdLineOpts.GetStatements(), ";")

	if cmdLineOpts.GetOutputFile() != "" {
		if cmdLineOpts.GetPretty() {
			displayOpts := &outputs.PrettyCSVOutputOptions{
				WriteHeader: cmdLineOpts.GetOutputHeader(),
				WriteTo:     util.OpenFileOrStdDev(cmdLineOpts.GetOutputFile(), true),
			}

			outputer = outputs.NewPrettyCSVOutput(displayOpts)
		} else {
			displayOpts := &outputs.CSVOutputOptions{
				WriteHeader: cmdLineOpts.GetOutputHeader(),
				Seperator:   util.DetermineSeparator(cmdLineOpts.GetOutputDelimiter()),
				WriteTo:     util.OpenFileOrStdDev(cmdLineOpts.GetOutputFile(), true),
			}

			outputer = outputs.NewCSVOutput(displayOpts)
		}
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
			saveErr := storage.SaveTo(tempFile.Name())

			if saveErr != nil {
				log.Fatalln(saveErr)
			}

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
