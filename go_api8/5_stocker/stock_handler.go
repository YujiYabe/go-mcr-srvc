package stocker

import (
	"github.com/jinzhu/gorm"
	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// SQLHandler ...
type SQLHandler struct {
	Conn *gorm.DB
}

// NewMySQLHandler ...
func NewMySQLHandler() *SQLHandler {
	conn, err := gorm.Open("mysql", "user:user@tcp(mysql)/app?charset=utf8&parseTime=True&loc=Local")
	conn.LogMode(true)

	if err != nil {
		panic(err.Error)
	}
	SQLHandler := new(SQLHandler)
	SQLHandler.Conn = conn
	return SQLHandler
}

// StockFind ...
func (handler *SQLHandler) StockFind(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
}

// StockFindByName ...
func (handler *SQLHandler) StockFindByName(out interface{}, where ...interface{}) *gorm.DB {

	// return handler.Conn.Where("name = ?", name).First(out)
	return handler.Conn.First(out, where...)
}

// StockFindByNames ...
// func (handler *SQLHandler) StockFindByNames(out interface{}, where ...interface{}) *gorm.DB {
// func (handler *SQLHandler) StockFindByNames(out interface{}, where []string) *gorm.DB {
func (handler *SQLHandler) StockFindByNames(where map[string]int) *gorm.DB {
	for item, num := range where {
		handler.Conn.Table("vegetables").Where("name IN (?)", item).UpdateColumn("stock", gorm.Expr("stock - ?", num))
	}
	return handler.Conn
}

// // INFRExec ...
// func (handler *SQLHandler) INFRExec(sql string, values ...interface{}) *gorm.DB {
// 	return handler.Conn.Exec(sql, values...)
// }

// INFRFirst ...
func (handler *SQLHandler) INFRFirst(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.First(out, where...)
}

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
