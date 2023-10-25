package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/OneCloudInc/textql/inputs"
	"github.com/OneCloudInc/textql/outputs"
	"github.com/OneCloudInc/textql/storage"
	"github.com/OneCloudInc/textql/util"
)

type commandLineOptions struct {
	Statements      *string
	StatementsFile  *string
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
	cmdLineOpts.StatementsFile = flag.String("sqlfile", "", "SQL filepath to run on the data")
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

func (clo *commandLineOptions) GetStatements() (string, *error) {
	if clo.Statements == nil && clo.StatementsFile == nil {
		err := fmt.Errorf("No SQL statements provided")
		return "", &err
	}
	if clo.Statements != nil && *clo.Statements != "" {
		return *clo.Statements, nil
	}
	filepath := *clo.StatementsFile
	dat, err := os.ReadFile(filepath)
	if err != nil {
		return "", &err
	}
	return string(dat), nil
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

func handleSemiColon(sqlStrings *[]string) []string {
	var stack []string
	var current string
	index := 0
	current = (*sqlStrings)[index]
	for {
		if current == "" {
			current = (*sqlStrings)[index]
		}
		count := strings.Count(current, "'")
		if count%2 == 1 {
			current = current + ";" + (*sqlStrings)[index+1]
		} else {
			stack = append(stack, current)
			current = ""
		}
		index++
		if index == len((*sqlStrings)) {
			break
		}
	}
	return stack
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

	for _, taggedName := range cmdLineOpts.GetSourceFiles() {
		// support <tablename>:<filename> syntax
		var names = strings.SplitN(taggedName, ":", 2)
		var sourceFile = names[len(names)-1]

		if util.IsPathDir(sourceFile) {
			for _, file := range util.AllFilesInDirectory(sourceFile) {
				inputSources = append(inputSources, file)
			}
		} else {
			inputSources = append(inputSources, taggedName)
		}
	}

	storage := storage.NewSQLite3StorageWithDefaults()

	for _, taggedName := range inputSources {
		// support <tablename>:<filename> syntax
		var names = strings.SplitN(taggedName, ":", 2)
		var file = names[len(names)-1]

		fp := util.OpenFileOrStdDev(file, false)

		inputOpts := &inputs.CSVInputOptions{
			HasHeader: cmdLineOpts.GetHeader(),
			Separator: util.DetermineSeparator(cmdLineOpts.GetDelimiter()),
			ReadFrom:  fp,
		}

		input, inputErr := inputs.NewCSVInput(inputOpts)

		if inputErr != nil {
			log.Printf("Unable to load %v\n", file)
		}

		if len(names) > 1 {
			input.SetName(names[0])
		}
		storage.LoadInput(input)
	}

	stat, err := cmdLineOpts.GetStatements()
	if err != nil {
		log.Fatalln(err)
	}
	if (strings.Count(stat, "'") % 2) == 1 {
		log.Fatalln("String contains odd number of \"'(Single Quotes)\"")
	}

	sqlStrings := strings.Split(stat, ";")

	sqlStrings = handleSemiColon(&sqlStrings)

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
				Separator:   util.DetermineSeparator(cmdLineOpts.GetOutputDelimiter()),
				WriteTo:     util.OpenFileOrStdDev(cmdLineOpts.GetOutputFile(), true),
			}

			outputer = outputs.NewCSVOutput(displayOpts)
		}
	}

	for _, sqlQuery := range sqlStrings {
		queryResults, queryErr := storage.ExecuteSQLString(sqlQuery)

		if queryErr != nil {
			log.Fatalln(queryErr)
		}

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
