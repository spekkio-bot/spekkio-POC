package sqlbuilder

import (
	"errors"
)

// SelectQueryProps structures properties of a SQL SELECT query to be built
type SelectQueryProps struct {
	BaseTable string
	Columns   []Column
}

// Column structures a table column name and alias for SelectQueryProps
// Alias can be left empty to use no alias
type Column struct {
	Name  string
	Alias string
}

// BuildQuery builds a SQL query from a SelectQueryProps object
func (q *SelectQueryProps) BuildQuery() (string, error) {
	sql := "SELECT "
	for i, column := range q.Columns {
		if i > 0 {
			sql += ", "
		}
		if column.Name == "" {
			return "", errors.New("empty Column.name not allowed")
		}
		sql += column.Name
		if column.Alias != "" {
			sql += " AS " + column.Alias
		}
	}
	sql += " FROM " + q.BaseTable
	sql += ";"
	return sql, nil
}
