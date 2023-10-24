package domain

type (
	Stock struct {
		AllProductList
		ProductList
		IsVaildJANCodeList []int
	}

	AllProductList []Product
	ProductList    []Product
)

func NewStock() *Stock {
	//ここでDBから取得する

	return &Stock{}
}

func NewProduct() Product {
	return Product{}
}

func NewProductList() ProductList {
	return ProductList{}
}

func NewAllProductList() ProductList {
	return ProductList{}
}

type (
	Product struct {
		JANCode       int    `db:"jan_code" json:"jan_code"`             // jan_code
		IsValid       bool   `db:"is_valid" json:"is_valid"`             // 販売終了の商品を無効にする
		Place         string `db:"place" json:"place"`                   // 冷蔵庫の場所
		NameJa        string `db:"name_ja" json:"name_ja"`               // 名前_ja
		NameEn        string `db:"name_en" json:"name_en"`               // 名前_en
		NameEs        string `db:"name_es" json:"name_es"`               // 名前_es
		NameZh        string `db:"name_zh" json:"name_zh"`               // 名前_zh
		NameAr        string `db:"name_ar" json:"name_ar"`               // 名前_ar
		InformationJa string `db:"information_ja" json:"information_ja"` // 商品のキャッチコピー_ja
		InformationEn string `db:"information_en" json:"information_en"` // 商品のキャッチコピー_en
		InformationEs string `db:"information_es" json:"information_es"` // 商品のキャッチコピー_es
		InformationZh string `db:"information_zh" json:"information_zh"` // 商品のキャッチコピー_zh
		InformationAr string `db:"information_ar" json:"information_ar"` // 商品のキャッチコピー_ar
		Recipe        string `db:"recipe" json:"recipe"`                 // レシピ
		CostPrice     int    `db:"cost_price" json:"cost_price"`         // 原価
		Calorie       int    `db:"calorie" json:"calorie"`               // カロリー
		RetailPrice   int    `db:"retail_price" json:"retail_price"`     // 小売価格
		OperationCost int    `db:"operation_cost" json:"operation_cost"` // 提供するまでかかった時間の価格
		EnergyCost    int    `db:"energy_cost" json:"energy_cost"`       // 提供するまでかかった電気料金
		Tag           string `db:"tag" json:"tag"`                       // タグ情報 json:array
		Allergy       string `db:"allergy" json:"allergy"`               // アレルギー情報 json:object
	}
)
