package domain

import (
	"context"

	"backend/pkg"
)

var (
	myErr *pkg.MyErr
)

// GetAllergyDefault ...
func (receiver *domain) GetAllergyDefault(
	ctx context.Context,
) Allergy {
	return receiver.AllergyDefault
}

// GetAllergyList ...
func (receiver *domain) GetAllergyList(
	ctx context.Context,
) AllergyList {
	return receiver.AllergyList
}

// GetIsVaildLangCodeMap ...
func (receiver *domain) GetIsVaildLangCodeMap() map[int]string {
	return receiver.isVaildLangCodeMap
}

func (receiver *domain) GetIsVaildLangCodeList() []int {
	return receiver.isVaildLangCodeList
}

func (receiver *domain) GetDefaultLangCode() int {
	return receiver.defaultLangCode
}

// SaveInMemory ...
func (receiver *domain) SaveInMemory(
	ctx context.Context,
	allProductList AllProductList,
) error {
	productList := ProductList{}
	isVaildJANCodeList := []int{}

	for _, allProduct := range allProductList {
		if allProduct.IsValid {
			productList = append(productList, allProduct)
			isVaildJANCodeList = append(isVaildJANCodeList, allProduct.JANCode)
		}
	}

	receiver.AllProductList = allProductList
	receiver.ProductList = productList
	receiver.IsVaildJANCodeList = isVaildJANCodeList

	return nil
}
