package supplier

import "github.com/jinzhu/gorm"

// VegetableSupplierToStocker ...
type VegetableSupplierToStocker interface {
	StockFind(interface{}, ...interface{}) *gorm.DB
	StockFindByName(interface{}, ...interface{}) *gorm.DB
	StockFindByNames(interface{}, ...interface{}) *gorm.DB
	// INFRExec(string, ...interface{}) *gorm.DB
	// INFRFirst(interface{}, ...interface{}) *gorm.DB
	// INFRRaw(string, ...interface{}) *gorm.DB
	// INFRCreate(interface{}) *gorm.DB
	// INFRSave(interface{}) *gorm.DB
	// INFRDelete(interface{}) *gorm.DB
	// INFRWhere(interface{}, ...interface{}) *gorm.DB
}
