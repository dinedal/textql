# TextQL

Allows you to easily execute SQL against structured text like CSV or TSV.

Example session:
![textql_usage_session](https://raw.github.com/dinedal/textql/master/textql_usage.gif)

## Is it any good?

[Yes](https://news.ycombinator.com/item?id=3067434)

## Requirements

- sqlite3
- Go

## Install

You may need to `export CC=clang` on OS X.

Install sqlite3 via whatever means you perfer, and:

```bash
go get -u github.com/dinedal/textql
```

## Usage

```bash
  -console=false: After all commands are run, open sqlit3 console with this data
  -dlm=",": Delimiter between fields, (\t for tab)
  -header=false: Treat file as having the first row as a header row
  -save-to="": If set, sqlite3 db is left on disk at this path
  -source="stdin": Source file to load, or defaults to stdin
  -sql="": SQL Command(s) to run on the data
  -table-name="tbl": Override the default table name (tbl)
  -verbose=false: Enable verbose logging
```
