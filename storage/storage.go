package storage

import (
	"database/sql"

	"github.com/dinedal/textql/inputs"
)

// Storage implentors are expected to be SQL capable engines.
// A Storage should support loading data from a TextQL input,
// executing any number of SQL statements and returning their resultant
// sql.Rows, as well as supporting clean closing and "backing up" of
// data to a specific path
type Storage interface {
	// LoadInput should make all the data in the input accessible to queries.
	LoadInput(*inputs.Input)
	// SaveTo should write the entire database of the SQL backend to the path given as a string.
	// Failure in any way should return an error, and nil if the operation was successful.
	SaveTo(string) error
	// ExecuteSQLString should first convert from TextQL shorthand SQL to normal SQL,
	// apply the query or transformation given to the SQL backend and return either nil
	// or the sql.Rows that were returned from the query.
	ExecuteSQLString(string) *sql.Rows
	// Close should cleanly close the database backend, cleaning up data on disk if required.
	Close()
}
