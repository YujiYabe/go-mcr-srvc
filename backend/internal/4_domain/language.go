package domain

func init() {
	for key := range isVaildlangCodeMap {
		isVaildlangCodeList = append(isVaildlangCodeList, key)
	}
}

const DefaultLangCode = 1041

var isVaildlangCodeList = []int{}

func GetIsVaildLangCodeList() []int {
	return isVaildlangCodeList
}

var isVaildlangCodeMap = map[int]string{
	1041:  "日本語",     // japanese
	1033:  "English", // us english
	1034:  "Español", // spanish
	2052:  "中文",      // 中国語 (中華人民共和国)
	14337: "عربي",    // アラビア語 (U.A.E.)
}

func GetIsVaildlangCodeMap() map[int]string {
	return isVaildlangCodeMap
}
