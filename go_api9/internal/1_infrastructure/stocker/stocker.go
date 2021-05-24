package stocker

import (
	"context"

	"github.com/jinzhu/gorm"
	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Stocker ...
type Stocker struct {
	Conn *gorm.DB
}

// NewToStocker ...
func NewToStocker() *Stocker {
	conn, err := gorm.Open("mysql", "user:user@tcp(mysql)/app?charset=utf8&parseTime=True&loc=Local")
	conn.LogMode(true)

	if err != nil {
		panic(err.Error)
	}
	s := new(Stocker)
	s.Conn = conn
	return s
}

// Dummy ...
func (s *Stocker) Dummy(ctx context.Context) error {
	return nil
}

// StockFind ...
func (s *Stocker) StockFind(out interface{}, where ...interface{}) (string, error) {
	_, err := s.Conn.Find(out, where...)
	if err != nil {
		return "", err
	}

	return "ok", nil
}
