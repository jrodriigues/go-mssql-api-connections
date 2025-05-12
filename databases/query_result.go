package databases

import "database/sql"

type QueryResult struct {
	Rows *sql.Rows
}
