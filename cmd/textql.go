package main

import (
	"flag"
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
	TableName       *string
	SaveTo          *string
	Console         *bool
}

func NewCommandLineOptions() *CommandLineOptions {
	opts := CommandLineOptions{}
	opts.Commands = flag.String("sql", "", "SQL Command(s) to run on the data")
	opts.SourceFile = flag.String("source", "stdin", "Source file to load, or defaults to stdin")
	opts.Delimiter = flag.String("dlm", ",", "Input delimiter between fields -dlm=tab for tab, -dlm=opts.0x## to specify a character code in hex")
	opts.Header = flag.Bool("header", false, "Treat file as having the first row as a header row")
	opts.OutputHeader = flag.Bool("output-header", false, "Display column names in output")
	opts.OutputDelimiter = flag.String("output-dlm", ",", "Output delimiter between fields -output-dlm=tab for tab, -dlm=0x## to specify a character code in hex")
	opts.OutputFile = flag.String("output-file", "stdout", "Filename to write output to, if empty no output is written")
	opts.TableName = flag.String("table-name", "", "Override the default table name (input file name or stdin)")
	opts.SaveTo = flag.String("save-to", "", "If set, sqlite3 db is left on disk at this path")
	opts.Console = flag.Bool("console", false, "After all commands are run, open sqlite3 console with this data")
	flag.Parse()

	return &opts
}

func (this *CommandLineOptions) GetCommands() string {
	return *this.Commands
}

func (this *CommandLineOptions) GetSourceFile() string {
	return *this.SourceFile
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

func (this *CommandLineOptions) GetTableName() string {
	return *this.TableName
}

func (this *CommandLineOptions) GetSaveTo() string {
	return *this.SaveTo
}

func (this *CommandLineOptions) GetConsole() bool {
	return *this.Console
}

func main() {
	cmdopts := NewCommandLineOptions()

	if cmdopts.GetConsole() {
		if cmdopts.GetSourceFile() == "stdin" {
			log.Fatalln("Can not open console with pipe input, read a file instead")
		}
		_, sqlite3ConsolePathErr := exec.LookPath("sqlite3")
		if sqlite3ConsolePathErr != nil {
			log.Fatalln("Console requested but unable to locate `sqlite3` program on $PATH")
		}
	}

	fp := util.OpenFileOrStdDev(cmdopts.GetSourceFile())

	opts := &inputs.CSVInputOptions{
		HasHeader: cmdopts.GetHeader(),
		Seperator: util.DetermineSeparator(cmdopts.GetDelimiter()),
		ReadFrom:  fp,
	}

	input := inputs.NewCSVInput(opts)

	storage_opts := &storage.SQLite3Options{}

	storage := storage.NewSQLite3Storage(storage_opts)

	if (cmdopts.GetTableName()) != "" {
		input.SetName(cmdopts.GetTableName())
	}

	storage.LoadInput(input)

	queryResults := storage.ExecuteSQLStrings(strings.Split(cmdopts.GetCommands(), ";"))

	if cmdopts.GetOutputFile() != "" {
		displayOpts := &outputs.CSVOutputOptions{
			WriteHeader: cmdopts.GetOutputHeader(),
			Seperator:   util.DetermineSeparator(cmdopts.GetOutputDelimiter()),
			WriteTo:     util.OpenFileOrStdDev(cmdopts.GetOutputFile()),
		}

		outputer := outputs.NewCSVOutput(displayOpts)
		outputer.Show(queryResults)
	}

	if cmdopts.GetSaveTo() != "" {
		storage.SaveTo(util.CleanPath(cmdopts.GetSaveTo()))
	}

	if cmdopts.GetConsole() {
		var args []string

		if cmdopts.GetOutputHeader() {
			args = []string{"-header"}
		} else {
			args = []string{}
		}

		if cmdopts.GetSaveTo() != "" {
			args = append(args, util.CleanPath(cmdopts.GetSaveTo()))
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
