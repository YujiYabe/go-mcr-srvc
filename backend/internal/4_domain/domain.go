package domain

import (
	"context"

	"backend/pkg"
)

var (
	myErr *pkg.MyErr
)

// ParseOrder ...
func (receiver *domain) ParseOrder(
	ctx context.Context,
) error {
	return nil
}

// CookHamburgers ...
func (receiver *domain) CookHamburgers(
	ctx context.Context,
) error {
	return nil
}

// GetAllergyDefault ...
func (receiver *domain) GetAllergyDefault(
	ctx *context.Context,
) *Allergy {
	return (*Allergy)(receiver.AllergyDefault)
}

// GetIsVaildLangCodeMap ...
func (receiver *domain) GetIsVaildLangCodeMap(
	ctx *context.Context,
) map[int]string {
	return receiver.isVaildLangCodeMap
}

// CookHamburgers ...
func (receiver *domain) SaveInMemory(
	ctx *context.Context,
	allProductList *AllProductList,
) error {
	productList := ProductList{}
	isVaildJANCodeList := []int{}

	for _, allProduct := range *allProductList {
		if allProduct.IsValid {
			productList = append(productList, allProduct)
			isVaildJANCodeList = append(isVaildJANCodeList, allProduct.JANCode)
		}
	}

	receiver.AllProductList = *allProductList
	receiver.ProductList = productList
	receiver.IsVaildJANCodeList = isVaildJANCodeList

	return nil
}
