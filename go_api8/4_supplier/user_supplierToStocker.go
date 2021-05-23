package supplier

import "github.com/jinzhu/gorm"

// UserSupplierToStocker ...
type UserSupplierToStocker interface {
	StockFind(interface{}, ...interface{}) *gorm.DB
	// INFRExec(string, ...interface{}) *gorm.DB
	// INFRFirst(interface{}, ...interface{}) *gorm.DB
	// INFRRaw(string, ...interface{}) *gorm.DB
	// INFRCreate(interface{}) *gorm.DB
	// INFRSave(interface{}) *gorm.DB
	// INFRDelete(interface{}) *gorm.DB
	// INFRWhere(interface{}, ...interface{}) *gorm.DB
}
