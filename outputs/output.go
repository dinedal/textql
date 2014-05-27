package outputs

import "database/sql"

type Output interface {
	Show([]*sql.Rows)
}
