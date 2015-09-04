package inputs

// Input is how TextQL reads from data sources.
// To be an input, an implementor must return tabular data.
// How data is manipulated into the tabular structure is left to the implementor.
// Inputs are expected to return in a row by row fashion.
//
// Returning a nil on ReadRecord() signifies that no more data will be returned
// because either the end of stream has been reached, or an error has
// occured that makes further reading impossible.
// Recoverable / ignoreable errors should represent themselves as empty sets.
//
// Header() should return metadata naming the columns in the table.
// No metadata on the columns or rows types is expected or supported.
//
// Name() should return a reasonable name for the data set, prehaps the file name
// or source data name, by default.
//
// SetName() allows users of the dataset to supply their own name if needed.
// Inputs should respect the new name if given one.
type Input interface {
	ReadRecord() []string
	Header() []string
	Name() string
	SetName(string)
}
