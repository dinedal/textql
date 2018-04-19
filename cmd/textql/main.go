package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/ivotron/textql"
)

func main() {
	// Parse command line opts
	commands := flag.String("sql", "", "SQL Command(s) to run on the data")
	source_text := flag.String("source", "stdin", "Source file to load, or defaults to stdin")
	delimiter := flag.String("dlm", ",", "Delimiter between fields -dlm=tab for tab, -dlm=0x## to specify a character code in hex")
	lazyQuotes := flag.Bool("lazy-quotes", false, "Enable LazyQuotes in the csv parser")
	header := flag.Bool("header", false, "Treat file as having the first row as a header row")
	outputHeader := flag.Bool("output-header", false, "Display column names in output")
	tableName := flag.String("table-name", "tbl", "Override the default table name (tbl)")
	save_to := flag.String("save-to", "", "If set, sqlite3 db is left on disk at this path")
	console := flag.Bool("console", false, "After all commands are run, open sqlite3 console with this data")
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	if *console && (*source_text == "stdin") {
		log.Fatalln("Can not open console with pipe input, read a file instead")
	}

	db, openPath, separator := textql.Load(
		source_text, delimiter, lazyQuotes, header,
		tableName, save_to, console, verbose)

	textql.Execute(db, commands, separator, verbose, outputHeader)

	// Open console
	if *console {
		db.Close()
		args := []string{*openPath}
		if *outputHeader {
			args = append(args, "-header")
		}
		cmd := exec.Command("sqlite3", args...)

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd_err := cmd.Run()
		if cmd.Process != nil {
			cmd.Process.Release()
		}

		if len(*save_to) == 0 {
			os.RemoveAll(filepath.Dir(*openPath))
		}

		if cmd_err != nil {
			log.Fatalln(cmd_err)
		}
	} else if len(*save_to) == 0 {
		db.Close()
		os.Remove(*openPath)
	} else {
		db.Close()
	}
}
