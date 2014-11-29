package main

const APP_NAME = "textql"
const APP_DESCRIPTION = "Execute SQL against structured text like CSV or TSV"

//
// command line flags
//

const FLAG_QUERIES = "sql"
const FLAG_QUERIES_HELP = "SQL command(s) to execute against the data"

const FLAG_INPUT = "input"
const FLAG_INPUT_HELP = "Input CSV file, or defaults to stdin"

const FLAG_DELIMITER = "dlm"
const FLAG_DELIMITER_HELP = "Delimiter between fields (tab, 0x## for hex specification)"

const FLAG_LAZY_QUOTES = "lazy-quotes"
const FLAG_LAZY_QUOTES_HELP = "Enable lazy quotes in the CSV parser"

const FLAG_HEADER = "header"
const FLAG_HEADER_HELP = "Treat the file as having the first row as a header row"

const FLAG_OUTPUT_HEADER = "output-header"
const FLAG_OUTPUT_HEADER_HELP = "Display column names in output"

const FLAG_TABLE_NAME = "table-name"
const FLAG_TABLE_NAME_HELP = "Override the default table name (tbl)"

const FLAG_SAVE_TO = "save-to"
const FLAG_SAVE_TO_HELP = "If set, sqlite3 DB is left on disk at this path"

const FLAG_CONSOLE = "console"
const FLAG_CONSOLE_HELP = "After all commands are run, open sqlite3 console with this data"

const FLAG_VERBOSE = "verbose"
const FLAG_VERBOSE_HELP = "Enable verbose logging"

//
// errors
//

const ERR_NO_CONSOLE_STDIN = "Can not open console with pipe input, read a file instead"

const DEFAULT_TABLE_NAME = "tbl"
