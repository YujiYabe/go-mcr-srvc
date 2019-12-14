package infrastructure

import (
	"github.com/jinzhu/gorm"

	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"app/interfaces/database"
)

// SQLHandler ...
type SQLHandler struct {
	Conn *gorm.DB
}

// NewSQLHandler ...
func NewSQLHandler() database.IFDBSQLHandler {
	conn, err := gorm.Open("mysql", "user:user@tcp(mysql)/app?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error)
	}
	SQLHandler := new(SQLHandler)
	SQLHandler.Conn = conn
	return SQLHandler
}

// Find ...
func (handler *SQLHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
}

// Exec ...
func (handler *SQLHandler) Exec(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Exec(sql, values...)
}

// First ...
func (handler *SQLHandler) First(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.First(out, where...)
}

// Raw ...
func (handler *SQLHandler) Raw(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Raw(sql, values...)
}

// Create ...
func (handler *SQLHandler) Create(value interface{}) *gorm.DB {
	return handler.Conn.Create(value)
}

// Save ...
func (handler *SQLHandler) Save(value interface{}) *gorm.DB {
	return handler.Conn.Save(value)
}

// Delete ...
func (handler *SQLHandler) Delete(value interface{}) *gorm.DB {
	return handler.Conn.Delete(value)
}

// Where ...
func (handler *SQLHandler) Where(query interface{}, args ...interface{}) *gorm.DB {
	return handler.Conn.Where(query, args...)
}
