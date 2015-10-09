package outputs

import "database/sql"

// Output implementors should accept sql.Rows and transform them
// however they need to in order to represent them in their specific format.
type Output interface {
	// Show should display/write the sql.Rows to the implmentor's destination and format.
	Show(*sql.Rows)
}
