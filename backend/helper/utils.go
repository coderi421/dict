package helper

import "unicode"

// AnalyzeInputType 同上
func AnalyzeInputType(s string) string {
	hasChinese := false
	hasEnglish := false
	//hasOther := false
	for _, r := range s {
		if unicode.Is(unicode.Han, r) {
			hasChinese = true
		} else if unicode.Is(unicode.Latin, r) {
			hasEnglish = true
		} else {
			//hasOther = true
		}
	}
	switch {
	case hasChinese && hasEnglish:
		return "mixed"
	case hasChinese:
		return "pure_chinese"
	case hasEnglish:
		return "pure_english"
	default:
		return "other"
	}
	//switch {
	//case hasChinese && !hasEnglish && !hasOther:
	//	return "pure_chinese"
	//case !hasChinese && hasEnglish && !hasOther:
	//	return "pure_english"
	//case hasChinese && hasEnglish && !hasOther:
	//	return "mixed"
	//default:
	//	return "other"
	//}
}
