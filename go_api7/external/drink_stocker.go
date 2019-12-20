package external

import (
	"github.com/jinzhu/gorm"

	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"app/interfaces/supplier"
)

// DrinkStocker ...
type DrinkStocker struct {
	Conn *gorm.DB
}

// NewDrinkStocker ...
func NewDrinkStocker() supplier.ExtractDrink {
	conn, err := gorm.Open("mysql", "user:user@tcp(mysql)/app?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error)
	}
	DrinkStocker := new(DrinkStocker)
	DrinkStocker.Conn = conn
	return DrinkStocker
}

// ExtractCoffee ...
func (handler *DrinkStocker) ExtractCoffee(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
}

// // INFRExec ...
// func (handler *DrinkStocker) INFRExec(sql string, values ...interface{}) *gorm.DB {
// 	return handler.Conn.Exec(sql, values...)
// }

// // INFRFirst ...
// func (handler *DrinkStocker) INFRFirst(out interface{}, where ...interface{}) *gorm.DB {
// 	return handler.Conn.First(out, where...)
// }

// // INFRRaw ...
// func (handler *DrinkStocker) INFRRaw(sql string, values ...interface{}) *gorm.DB {
// 	return handler.Conn.Raw(sql, values...)
// }

// // INFRCreate ...
// func (handler *DrinkStocker) INFRCreate(value interface{}) *gorm.DB {
// 	return handler.Conn.Create(value)
// }

// // INFRSave ...
// func (handler *DrinkStocker) INFRSave(value interface{}) *gorm.DB {

// 	return handler.Conn.Save(value)
// }

// // INFRDelete ...
// func (handler *DrinkStocker) INFRDelete(value interface{}) *gorm.DB {
// 	return handler.Conn.Delete(value)
// }

// // INFRWhere ...
// func (handler *DrinkStocker) INFRWhere(query interface{}, args ...interface{}) *gorm.DB {
// 	return handler.Conn.Where(query, args...)
// }
