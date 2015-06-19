package storage

import (
	"database/sql"

	"github.com/dinedal/textql/inputs"
)

type Storage interface {
	LoadInput(*inputs.Input)
	SaveTo(string)
	ExecuteSQLString(string) *sql.Rows
	Close()
}
