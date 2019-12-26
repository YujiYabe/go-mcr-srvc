package db

import (
	"github.com/jinzhu/gorm"

	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// SQLHandler ...
type SQLHandler struct {
	Conn *gorm.DB
}

// NewSQLHandler ...
func NewSQLHandler() *SQLHandler {
	conn, err := gorm.Open("mysql", "user:user@tcp(mysql)/app?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error)
	}
	SQLHandler := new(SQLHandler)
	SQLHandler.Conn = conn
	return SQLHandler
}

// INFRFind ...
func (handler *SQLHandler) INFRFind(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
}

// // INFRExec ...
// func (handler *SQLHandler) INFRExec(sql string, values ...interface{}) *gorm.DB {
// 	return handler.Conn.Exec(sql, values...)
// }

// // INFRFirst ...
// func (handler *SQLHandler) INFRFirst(out interface{}, where ...interface{}) *gorm.DB {
// 	return handler.Conn.First(out, where...)
// }

// // INFRRaw ...
// func (handler *SQLHandler) INFRRaw(sql string, values ...interface{}) *gorm.DB {
// 	return handler.Conn.Raw(sql, values...)
// }

// // INFRCreate ...
// func (handler *SQLHandler) INFRCreate(value interface{}) *gorm.DB {
// 	return handler.Conn.Create(value)
// }

// // INFRSave ...
// func (handler *SQLHandler) INFRSave(value interface{}) *gorm.DB {

// 	return handler.Conn.Save(value)
// }

// // INFRDelete ...
// func (handler *SQLHandler) INFRDelete(value interface{}) *gorm.DB {
// 	return handler.Conn.Delete(value)
// }

// // INFRWhere ...
// func (handler *SQLHandler) INFRWhere(query interface{}, args ...interface{}) *gorm.DB {
// 	return handler.Conn.Where(query, args...)
// }
