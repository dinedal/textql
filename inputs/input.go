package inputs

// Input is how TextQL reads from data sources.
// To be an input, an implementor must return tabular data.
// How data is manipulated into the tabular structure is left to the implementor.
// Inputs are expected to return in a row by row fashion.
type Input interface {
	// ReadRecord should return nil on the end of data, or a single record.
	// Recoverable errors should represent themselves as empty sets.
	// Unrecoverable errors should return nil.
	ReadRecord() []string
	// Header should return metadata naming the columns in the table.
	Header() []string
	// Name should return a reasonable name for the data set, prehaps the file name.
	Name() string
	// SetName allows users of the dataset to supply their own name if needed.
	SetName(string)
}
