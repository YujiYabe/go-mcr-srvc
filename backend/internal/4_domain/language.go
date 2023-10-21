package domain

type (
	Language struct {
		defaultLangCode     int
		isVaildLangCodeList []int
		isVaildLangCodeMap  map[int]string
	}
)

func NewLanguage() *Language {
	language := &Language{
		defaultLangCode: 1041,
	}

	language.isVaildLangCodeMap = map[int]string{
		1041:  "日本語",     // japanese
		1033:  "English", // us english
		1034:  "Español", // spanish
		2052:  "中文",      // 中国語 (中華人民共和国)
		14337: "عربي",    // アラビア語 (U.A.E.)
	}

	var langList = []int{}
	for key := range language.isVaildLangCodeMap {
		langList = append(langList, key)
	}
	language.isVaildLangCodeList = langList

	return language
}

func (receiver *Language) GetIsVaildLangCodeList() []int {
	return receiver.isVaildLangCodeList
}

func (receiver *Language) GetIsVaildlangCodeMap() map[int]string {
	return receiver.isVaildLangCodeMap
}
