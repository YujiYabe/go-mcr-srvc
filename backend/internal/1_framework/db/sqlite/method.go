package sqlite

import (
	domain "backend/internal/4_domain"
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// UpdatePatties ...
func (s *Sqlite) UpdatePatties(ctx context.Context, items map[string]int) error {
	for item, num := range items {
		res := s.Conn.
			Table("patties").
			Where("name IN (?)", item).
			UpdateColumn("stock", gorm.Expr("stock - ?", num))

		if res.Error != nil {
			myErr.Logging(res.Error)
			return res.Error
		}

		// 作業時間を擬似的に再現
		time.Sleep(1 * time.Second)
	}
	return nil
}

// GetAllProductList ...
func (s *Sqlite) GetAllProductList(ctx *context.Context) *domain.AllProductList {
	allProductList := &domain.AllProductList{}

	s.Conn.Find(allProductList)

	fmt.Println("== Sqlite == == == == == == == == == ")
	fmt.Printf("%#v\n", allProductList)
	fmt.Println("== == == == == == == == == == ")

	return allProductList
}
