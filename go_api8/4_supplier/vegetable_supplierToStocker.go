package supplier

import "github.com/jinzhu/gorm"

// VegetableSupplierToStocker ...
type VegetableSupplierToStocker interface {
	StockFind(interface{}, ...interface{}) *gorm.DB
	StockFindByName(interface{}, ...interface{}) *gorm.DB
	// StockFindByNames(interface{}, ...interface{}) *gorm.DB
	// StockFindByNames(interface{}, []string) *gorm.DB
	StockFindByNames(map[string]int) *gorm.DB
}
