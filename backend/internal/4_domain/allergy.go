package domain

type (
	Allergy        map[string]string
	AllergyList    []Allergy
	AllergyDefault Allergy
)

func NewAllergyList() *AllergyList {
	return &AllergyList{
		newAllergyListJa(),
		newAllergyListEn(),
		newAllergyListEs(),
		newAllergyListZh(),
		newAllergyListAr(),
	}
}

func NeAllergyDefault() Allergy {
	return newAllergyListJa()
}

// https://www.relief.jp/docs/001403.html
func newAllergyListJa() Allergy {
	return Allergy{
		"lang_type": "1041", // japanese

		// 乳卵
		"egg":  "卵",
		"milk": "乳",

		// 魚介
		"shrimp":     "えび",
		"crab":       "かに",
		"squid":      "いか",
		"salmon":     "さけ",
		"mackerel":   "さば",
		"salmon_roe": "いくら",
		"abalone":    "あわび",

		// 種子
		"peanut":     "ピーナッツ",
		"walnut":     "くるみ",
		"soybean":    "大豆",
		"almond":     "アーモンド",
		"cashew_nut": "カシューナッツ",
		"sesame":     "ごま",

		// 穀物
		"wheat":     "小麦",
		"buckwheat": "そば",

		// フルーツ
		"apple":      "りんご",
		"orange":     "オレンジ",
		"banana":     "バナナ",
		"peach":      "もも",
		"kiwi_fruit": "キウイフルーツ",

		// 肉
		"pork":    "豚肉",
		"chicken": "鶏肉",
		"beef":    "牛肉",

		// その他
		"wild_yam":           "やまいも",
		"matsutake_mushroom": "まつたけ",
		"gelatin":            "ゼラチン",
	}
}

func newAllergyListEn() Allergy {
	return Allergy{
		"lang_type": "1033", // us english
		// 乳卵
		"egg":  "Egg",
		"milk": "Milk",

		// 魚介
		"shrimp":     "Shrimp",
		"crab":       "Crab",
		"squid":      "Squid",
		"salmon":     "Salmon",
		"mackerel":   "Mackerel",
		"salmon_roe": "Salmon roe",
		"abalone":    "Abalone",

		// 種子
		"peanut":     "Peanut",
		"walnut":     "Walnut",
		"soybean":    "Soybean",
		"almond":     "Almond",
		"cashew_nut": "Cashew nut",
		"sesame":     "Sesame",

		// 穀物
		"wheat":     "Wheat",
		"buckwheat": "Buckwheat",

		// フルーツ
		"apple":      "Apple",
		"orange":     "Orange",
		"banana":     "Banana",
		"peach":      "Peach",
		"kiwi_fruit": "Kiwi fruit",

		// 肉
		"pork":    "Pork",
		"chicken": "Chicken",
		"beef":    "Beef",

		// その他
		"wild_yam":           "Wild yam",
		"matsutake_mushroom": "Matsutake mushroom",
		"gelatin":            "Gelatin",
	}
}

func newAllergyListEs() Allergy {
	return Allergy{
		"lang_type": "1034", // spanish

		// 乳卵
		"egg":  "Huevo",
		"milk": "Leche",

		// 魚介
		"shrimp":     "Camarón",
		"crab":       "Cangrejo",
		"squid":      "Calamar",
		"salmon":     "Salmón",
		"mackerel":   "Caballa",
		"salmon_roe": "Huevas de salmón",
		"abalone":    "Abulón",

		// 種子
		"peanut":     "Maní",
		"walnut":     "Nuez",
		"soybean":    "Soja",
		"almond":     "Almendra",
		"cashew_nut": "Anacardo",
		"sesame":     "Sésamo",

		// 穀物
		"wheat":     "Trigo",
		"buckwheat": "Trigo sarraceno",

		// フルーツ
		"apple":      "Manzana",
		"orange":     "Naranja",
		"banana":     "Plátano",
		"peach":      "Durazno",
		"kiwi_fruit": "Kiwi",

		// 肉
		"pork":    "Cerdo",
		"chicken": "Pollo",
		"beef":    "Carne de res",

		// その他
		"wild_yam":           "Ñame silvestre",
		"matsutake_mushroom": "Seta matsutake",
		"gelatin":            "Gelatina",
	}
}

func newAllergyListZh() Allergy {
	return Allergy{
		"lang_type": "2052", // 中国語 (中華人民共和国)

		// 乳卵
		"egg":  "鸡蛋",
		"milk": "牛奶",

		// 魚介
		"shrimp":     "虾",
		"crab":       "蟹",
		"squid":      "鱿鱼",
		"salmon":     "鲑鱼",
		"mackerel":   "鲭鱼",
		"salmon_roe": "鲑鱼籽",
		"abalone":    "鲍鱼",

		// 種子
		"peanut":     "花生",
		"walnut":     "核桃",
		"soybean":    "大豆",
		"almond":     "杏仁",
		"cashew_nut": "腰果",
		"sesame":     "芝麻",

		// 穀物
		"wheat":     "小麦",
		"buckwheat": "荞麦",

		// フルーツ
		"apple":      "苹果",
		"orange":     "橙子",
		"banana":     "香蕉",
		"peach":      "桃子",
		"kiwi_fruit": "猕猴桃",

		// 肉
		"pork":    "猪肉",
		"chicken": "鸡肉",
		"beef":    "牛肉",

		// その他
		"wild_yam":           "山药",
		"matsutake_mushroom": "松茸",
		"gelatin":            "明胶",
	}
}

func newAllergyListAr() Allergy {
	return Allergy{
		"lang_type": "14337", // アラビア語 (U.A.E.)

		// 乳卵
		"egg":  "بيض",
		"milk": "حليب",

		// 魚介
		"shrimp":     "روبيان",
		"crab":       "سرطان البحر",
		"squid":      "حبار",
		"salmon":     "سمك السلمون",
		"mackerel":   "سمك الأسماك",
		"salmon_roe": "بيض السلمون",
		"abalone":    "أبلون",

		// 種子
		"peanut":     "فول السوداني",
		"walnut":     "جوز",
		"soybean":    "فول الصويا",
		"almond":     "لوز",
		"cashew_nut": "كاجو",
		"sesame":     "سمسم",

		// 穀物
		"wheat":     "قمح",
		"buckwheat": "الغرينية",

		// フルーツ
		"apple":      "تفاحة",
		"orange":     "البرتقال",
		"banana":     "موز",
		"peach":      "خوخ",
		"kiwi_fruit": "كيوي",

		// 肉
		"pork":    "لحم الخنزير",
		"chicken": "دجاج",
		"beef":    "لحم بقر",

		// その他
		"wild_yam":           "اليام البري",
		"matsutake_mushroom": "فطر ماتسوتاكي",
		"gelatin":            "جيلاتين",
	}
}
