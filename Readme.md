# TextQL

Allows you to easily execute SQL against structured text like CSV or TSV.

Example session:

```bash
~ cat ~/sample_data.csv
id,name,value
1,Paul,5
2,Jeff,16
3,Dmitri,-3

~ textql -source ~/sample_data.csv -sql "select count(*) from tbl;"
4
~ textql -header -source ~/sample_data.csv -sql "select sum(cast(value as INTEGER)) from tbl;"
18
~ cat ~/sample_data.csv | textql -header -sql "select max(id) from tbl;"
3
~ textql -header -source ~/sample_data.csv -console
SQLite version 3.7.13 2012-07-17 17:46:21
Enter ".help" for instructions
Enter SQL statements terminated with a ";"
sqlite> -- Full Fledged SQLite console
sqlite>
```

## Is it any good?

[Yes](https://news.ycombinator.com/item?id=3067434)

## Requirements

- sqlite3
- Go

## Install

You may need to `export CC=clang` on OS X.

Install sqlite3 via whatever means you perfer, and:

```bash
go get -u https://github.com/dinedal/textql
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
