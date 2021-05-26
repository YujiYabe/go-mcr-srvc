package stocker

import (
	"context"
	"fmt"
	"time"

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
	conn, err := open(30)
	if err != nil {
		panic(err)
	}

	s := new(Stocker)
	s.Conn = conn
	return s
}

func open(count uint) (*gorm.DB, error) {
	dsn := "user:user@tcp(mysql)/app?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		if count == 0 {
			return nil, fmt.Errorf("Retry count over")
		}
		time.Sleep(time.Second)
		fmt.Println(" ============================== ")
		fmt.Println(" wait db ...")
		fmt.Println(" ============================== ")

		// カウントダウンさせるようにする
		count--
		return open(count)
	}

	return db, nil
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
