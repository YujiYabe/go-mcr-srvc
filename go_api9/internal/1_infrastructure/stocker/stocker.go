package stocker

import (
	"context"
	// mysql
	// _ "github.com/jinzhu/gorm/dialects/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"app/internal/2_adapter/service"
)

// Stocker ...
type Stocker struct {
	Conn *gorm.DB
}

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
	// 	_, err := s.Conn.Find(out, where...)
	// 	if err != nil {
	// 		return "", err
	// 	}

	return "ok", nil
}
