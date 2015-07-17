# TextQL

Allows you to easily execute SQL against structured text like CSV or TSV.

Example session:
![textql_usage_session](https://raw.github.com/dinedal/textql/master/textql_usage.gif)

## Major changes!

Since there has been some time since the initial release of textql, I've made some improvements as well as made the project much more modular. There's also been a additional performance tweaks and added functionality, but this comes at the cost of breaking the original command line flags and changing the install command.

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

Bug fixes:

- Writing to a directory no longer fails silently.

## Key differences between textql and sqlite importing

- sqlite import will not accept stdin, breaking unix pipes. textql will happily do so.
- textql supports quote escaped delimiters, sqlite does not.
- textql leverages the sqlite in memory database feature as much as possible and only touches disk if asked.

## Is it any good?

[Yes](https://news.ycombinator.com/item?id=3067434)

## Requirements

- Go

## Install

You may need to `export CC=clang` on OS X.

```bash
go get -u github.com/dinedal/textql/...
```

## Usage

```bash
  textql [-console] [-save-to path] [-output-file path] [-output-dlm] [-output-header] [-header] [-dlm delimter] [-source path] [-sql sql_statements] [path ...]

  -console=false: After all commands are run, open sqlite3 console with this data
  -dlm=",": Input delimiter between fields -dlm=tab for tab, -dlm=opts.0x## to specify a character code in hex
  -header=false: Treat file as having the first row as a header row
  -output-dlm=",": Output delimiter between fields -output-dlm=tab for tab, -dlm=0x## to specify a character code in hex
  -output-file="stdout": Filename to write output to, if empty no output is written
  -output-header=false: Display column names in output
  -save-to="": If set, sqlite3 db is left on disk at this path
  -source="stdin": Source file or directory to load, or defaults to stdin
  -sql="": SQL Command(s) to run on the data
```


## License

New MIT License - Copyright (c) 2014, 2015 Paul Bergeron

See LICENSE for details
