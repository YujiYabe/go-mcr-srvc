package sqlite

import (
	"context"

	domain "backend/internal/4_domain"
)

// GetAllProductList ...
func (receiver *Sqlite) GetAllProductList(ctx context.Context) *domain.AllProductList {
	allProductList := &domain.AllProductList{}

	receiver.Conn.Find(allProductList)

	return allProductList
}

// UpdateProduct ...
func (receiver *Sqlite) UpdateProduct(
	ctx context.Context,
	newProduct domain.Product,
) {
	product := &domain.Product{}

	receiver.Conn.
		// Debug().
		Where("jan_code = ?", newProduct.JANCode).
		First(&product)

	product.IsValid = newProduct.IsValid
	product.Place = newProduct.Place

	receiver.Conn.
		// Debug().
		Where("jan_code = ?", newProduct.JANCode).
		Save(product)

}
