package sqlparser

// Magicify runs the SQL passed in, and a table name, throught a customized
// TextQL SQL Parser. This provides the following functionality:
//  - Queries that do not start with SELECT are implictly mapped to SELECT statements
//  - Queries that are missing a FROM, have the FROM inserted with tableName
func Magicify(sql string, tableName string) string {
	if tableName == "" {
		return sql
	}

	statement, err := Parse(sql)

	if err != nil {
		return sql
	}

	switch statement := statement.(type) {
	case *Select:
		if statement.From == nil {
			tableName := &TableName{[]byte(tableName), nil}
			aliasedTableExpr := AliasedTableExpr{tableName, nil, nil}
			tableExprs := TableExprs{&aliasedTableExpr}
			statement.From = &From{Type: AST_FROM, Expr: tableExprs}
		}
		return generateQuery(statement)
	default:
		return sql
	}
}

func generateQuery(statement Statement) string {
	buf := NewTrackedBuffer(nil)
	statement.Format(buf)
	return buf.String()
}
