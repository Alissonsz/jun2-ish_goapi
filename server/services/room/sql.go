package room

import _ "embed"

var (
	//go:embed sql/insert.sql
	insertQuery string
)
