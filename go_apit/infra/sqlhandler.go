package infra

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// SQLHandler ...
type SQLHandler struct {
	Conn *sql.DB
}

// NewSQLHandler ...
func NewSQLHandler() *SQLHandler {
	conn, err := sql.Open("mysql", "user:user@tcp(mysql:3306)/app")
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
