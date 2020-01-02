package stocker

import (
	"log"
	"runtime"

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
	debug := where
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	log.Println("====================================")
	log.Printf("%s:%d %s\n", file, line, f.Name())
	log.Printf("%v\n", debug)
	log.Println("------------------------------------")
	log.Printf("%+v\n", debug)
	log.Println("------------------------------------")
	log.Printf("%+v\n", debug)
	log.Println("------------------------------------")
	log.Printf("%#v\n", debug)
	log.Println("====================================")

	// return handler.Conn.Where("name = ?", name).First(out)
	return handler.Conn.First(out, where...)
}

// StockFindByNames ...
func (handler *SQLHandler) StockFindByNames(out interface{}, where ...interface{}) *gorm.DB {
	debug := where
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	log.Println("====================================")
	log.Printf("%s:%d %s\n", file, line, f.Name())
	log.Printf("%v\n", debug)
	log.Println("------------------------------------")
	log.Printf("%+v\n", debug)
	log.Println("------------------------------------")
	log.Printf("%+v\n", debug)
	log.Println("------------------------------------")
	log.Printf("%#v\n", debug)
	log.Println("====================================")

	// db.Where("name IN (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	// return handler.Conn.Where("name = ?", name).First(out)
	return handler.Conn.Where("name IN (?)", []string{"tomato", "lettuce"}).Find(out)
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
