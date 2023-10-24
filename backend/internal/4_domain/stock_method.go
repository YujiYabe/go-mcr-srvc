package domain

import (
	"context"
	"sort"

	"golang.org/x/exp/slices"
)

func (receiver *Stock) InitProduct(
	apl []Product,
) {
	receiver.FilterProductValid()
}

func (receiver *Stock) GetAllProductList() []Product {
	return receiver.AllProductList
}

func (receiver *Stock) GetProductList() []Product {
	return receiver.ProductList
}

func (receiver *Stock) GetProduct(
	ctx context.Context,
	janCode int,
) *Product {

	for _, product := range receiver.AllProductList {
		if janCode == product.JANCode {
			return &product
		}
	}
	return nil
}

// func GetProduct(janCode int) *Product {
// 	for _, product := range receiver.AllProductList {
// 		if janCode == product.JANCode {
// 			return &product
// 		}
// 	}
// 	return nil
// }

func (receiver *Stock) GetIsVaildJANCodeList() []int {
	return receiver.IsVaildJANCodeList
}

// UpdateProduct は指定された新製品の情報で、製品リスト中の該当製品を更新します。
func (receiver *Stock) UpdateProduct(newProduct *Product) {
	for index, product := range receiver.AllProductList {
		if newProduct.JANCode == product.JANCode {
			// インメモリ更新
			receiver.AllProductList[index].IsValid = newProduct.IsValid
			receiver.AllProductList[index].Place = newProduct.Place
			break
		}
	}
}

func (receiver *Stock) FilterProductValid() {
	receiver.ProductList = []Product{}
	receiver.IsVaildJANCodeList = []int{}

	for _, allProduct := range receiver.AllProductList {
		if allProduct.IsValid {
			receiver.ProductList = append(receiver.ProductList, allProduct)
			receiver.IsVaildJANCodeList = append(receiver.IsVaildJANCodeList, allProduct.JANCode)
		}
	}
}

func (receiver *Stock) VerifyJANCodes(
	JANCodeList []int,
	IsVaildJANCodeList []int,
	isVaildLangCodeList []int,
	defaultLangCode int,
) (
	[]int,
	int,
) {
	newIsVaildJANCodeList := []int{}
	newIsVaildLangCodeList := []int{}

	sort.Ints(JANCodeList)

	for _, janCode := range JANCodeList {
		if slices.Contains(isVaildLangCodeList, janCode) {
			newIsVaildLangCodeList = append(newIsVaildLangCodeList, janCode)
		}

		if slices.Contains(receiver.IsVaildJANCodeList, janCode) {
			newIsVaildJANCodeList = append(newIsVaildJANCodeList, janCode)
		}
	}

	newIsVaildLangCode := defaultLangCode // default 1041 japanese

	if len(newIsVaildLangCodeList) == 1 { // 言語が１つ指定されていればその言語にする。それ以外の場合変更しない。
		newIsVaildLangCode = newIsVaildLangCodeList[0]
	}

	return newIsVaildJANCodeList, newIsVaildLangCode
}
