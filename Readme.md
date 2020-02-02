# TextQL

[![Build Status](https://travis-ci.org/dinedal/textql.svg)](https://travis-ci.org/dinedal/textql) [![Go Report Card](http://goreportcard.com/badge/dinedal/textql)](http://goreportcard.com/report/dinedal/textql)

Allows you to easily execute SQL against structured text like CSV or TSV.

Example session:

![textql_usage_session](https://raw.github.com/dinedal/textql/master/textql_usage.gif)

## Major changes!

In the time since the initial release of textql, I've made some improvements as well as made the project much more modular. There've also been additional performance tweaks and added functionality, but this comes at the cost of breaking the original command-line flags and changing the install command.

### Changes since v1

Additions:

- Numeric values are automatically recognized in more cases.
- Date / Time / DateTime values are automatically recognized in reasonable formats. See [Time Strings](https://www.sqlite.org/lang_datefunc.html) for a list for accepted formats, and how to convert from other formats.
- Added join support! Multiple files / directories can be loaded by listing them at the end of the command.
- Directories are read by reading each file inside, and this is non-recursive.
- You can list as many files / directories as you like.
- Added flag '-output-file' to save output directly to a file.
- Added flag '-output-dlm' to modify the output delimiter.
- Added "short SQL" syntax.
  - For the case of a single table, the `FROM [table]` can be dropped from the query.
  - For simple selects, the `SELECT` keyword can be dropped from the query.
  - This means the v1 command `textql -sql "select * from tbl" -source some_file.csv` can be shortened to `textql -sql "*" some_file.csv`

Changes:

- The flag '-outputHeader' was renamed to '-output-header'.

Removals:

- Dropped the ability to override table names. This makes less sense after the automatic tablename generation based on filename, joins, and shorter SQL syntax changes.
- Removed '-source', any files / paths at the end of the command are used, as well as piped-in data.

Bug fixes:

- Writing to a directory no longer fails silently.

## Key differences between textql and sqlite importing

- sqlite import will not accept stdin, breaking unix pipes. textql will happily do so.
- textql supports quote-escaped delimiters, sqlite does not.
- textql leverages the sqlite in-memory database feature as much as possible and only touches disk if asked.

## Is it any good?

[Yes](https://news.ycombinator.com/item?id=3067434)

## Requirements

- Go 1.4 or later

## Install

**Latest release on Homebrew (OS X)**

```bash
brew install textql
```

**Build from source**

```bash
go get -u github.com/dinedal/textql/...
```

## Docker

First build the image.

```bash
docker build -t textql .
```

Now use that image mounting your current directory into the container.

```bash
docker run --rm -it -v $(pwd):/tmp textql [rest_of_command]
```

### Alias

You can add the following alias to your system to provide quick access to TextQL:

```bash
alias textql='docker run --rm -it -v $(pwd):/tmp textql '
```

## Usage

```bash
  textql [-console] [-save-to path path] [-output-file path] [-output-dlm delimter] [-output-header] [-pretty] [-quiet] [-header] [-dlm delimter] [-sql sql_statements] [path ...]

  -console
        After all statements are run, open SQLite3 REPL with this data
  -dlm string
        Input delimiter character between fields -dlm=tab for tab, -dlm=0x## to specify a character code in hex (default ",")
  -header
        Treat input files as having the first row as a header row
  -output-dlm string
        Output delimiter character between fields -output-dlm=tab for tab, -dlm=0x## to specify a character code in hex (default ",")
  -output-file file
        Filename to write output to, if empty no output is written (default "stdout")
  -output-header
        Display column names in output
  -quiet
        Suppress logging
  -pretty
        Pretty print output
  -save-to file
        SQLite3 db is left on disk at this file
  -sql string
        SQL Statement(s) to run on the data
  -version
        Print version and exit
```

## I want stdev, average, other functions

Just follow the install directions at [go-sqlite3-extension-functions](https://github.com/dinedal/go-sqlite3-extension-functions) and textql will automatically load this library.

Full function list:

- Math: acos, asin, atan, atn2, atan2, acosh, asinh, atanh, difference, degrees, radians, cos, sin, tan, cot, cosh, sinh, tanh, coth, exp, log, log10, power, sign, sqrt, square, ceil, floor, pi.
- String: replicate, charindex, leftstr, rightstr, ltrim, rtrim, trim, replace, reverse, proper, padl, padr, padc, strfilter.
- Aggregate: stdev, variance, mode, median, lower_quartile, upper_quartile

## License

New MIT License - Copyright (c) 2015, 2016 Paul Bergeron [http://pauldbergeron.com/](http://pauldbergeron.com/)

See LICENSE for details
