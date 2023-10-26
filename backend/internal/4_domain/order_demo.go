package domain

var (
	yakisoba = 4548779706595
	yakionig = 4902150657300
	pastaika = 4548779734116
	pastakan = 4548779734192
)

var (
	languageCodeJa = 1041  // japanese
	languageCodeEn = 1033  // us english
	languageCodeEs = 1034  // spanish
	languageCodeZh = 2052  // 中国語 (中華人民共和国)
	languageCodeAr = 14337 // アラビア語 (U.A.E.)
)

var ReservingListDemo = []Reserving{
	{
		QueueNo:      0,
		LanguageCode: languageCodeJa,
		JANCodeList: []int{
			yakisoba,
			yakisoba,
			pastakan,
			pastaika,
		},
	},
	{
		QueueNo:      1,
		LanguageCode: languageCodeEn,
		JANCodeList: []int{
			yakisoba,
			pastakan,
		},
	},
	{
		QueueNo:      2,
		LanguageCode: languageCodeEs,
		JANCodeList: []int{
			yakionig,
		},
	},
	{
		QueueNo:      3,
		LanguageCode: languageCodeZh,
		JANCodeList: []int{
			pastaika,
		},
	},
	{
		QueueNo:      4,
		LanguageCode: languageCodeAr,
		JANCodeList:  []int{},
	},
	{
		QueueNo:      10,
		LanguageCode: languageCodeAr,
		JANCodeList: []int{
			pastaika,
			pastakan,
		},
	},
}

var SoldListDemo = []Sold{
	{
		SoldNo:       1,
		LanguageCode: languageCodeAr,
		Status:       StatusPreparing,
		JANCodeList: []int{
			pastaika,
			pastaika,
			pastaika,
		},
	},
	{
		SoldNo:       2,
		LanguageCode: languageCodeJa,
		Status:       StatusCompleted,
		JANCodeList: []int{
			pastakan,
			pastaika,
			pastakan,
			yakionig,
		},
	},
	{
		SoldNo:       3,
		LanguageCode: languageCodeEn,
		Status:       StatusCompleted,
		JANCodeList: []int{
			yakisoba,
			pastaika,
		},
	},
	{
		SoldNo:       4,
		LanguageCode: languageCodeEs,
		Status:       StatusPreparing,
		JANCodeList: []int{
			yakisoba,
		},
	},
	{
		SoldNo:       5,
		LanguageCode: languageCodeZh,
		Status:       StatusCompleted,
		JANCodeList: []int{
			yakisoba,
			yakisoba,
			pastaika,
			pastakan,
		},
	},
	{
		SoldNo:       6,
		LanguageCode: languageCodeAr,
		Status:       StatusPassed,
		JANCodeList: []int{
			yakisoba,
		},
	},
	{
		SoldNo:       7,
		LanguageCode: languageCodeJa,
		Status:       StatusPreparing,
		JANCodeList: []int{
			pastaika,
			yakisoba,
		},
	},
	{
		SoldNo:       8,
		LanguageCode: languageCodeJa,
		Status:       StatusPassed,
		JANCodeList: []int{
			yakisoba,
			yakionig,
		},
	},
}
