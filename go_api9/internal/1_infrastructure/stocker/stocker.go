package stocker

import (
	"context"
	"fmt"

	// mysql
	// _ "github.com/jinzhu/gorm/dialects/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"app/internal/2_adapter/service"
)

type (
	// Stocker ...
	Stocker struct {
		Conn *gorm.DB
	}

	// Vegetables ...
	Vegetables struct {
		ID    int
		Name  string
		Stock int
	}
)

// NewToStocker ...
func NewToStocker() service.ToStocker {
	dsn := "user:user@tcp(mysql)/app?charset=utf8&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error)
	}
	// conn.LogMode(true)

	s := new(Stocker)
	s.Conn = conn
	return s
}

// Dummy ...
func (s *Stocker) Dummy(ctx context.Context) (string, error) {
	return "dummy ok", nil
}

// StockFind ...
func (s *Stocker) StockFind(out interface{}, where ...interface{}) (string, error) {
	vegetables := &Vegetables{}

	res := s.Conn.First(vegetables)
	if res.Error != nil {
		return "", res.Error
	}

	fmt.Println(" ============================== ")
	fmt.Printf("%+v\n", vegetables)
	fmt.Println(" ============================== ")

	return "ok", nil
}
