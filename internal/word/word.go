package word

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
	"unicode"
)

// ToUpper 小写单词到大写单词
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower 大写单词到小写单词
func ToLower(s string) string {
	return strings.ToLower(s)
}

// UnderscoreToUpperCamelCase 下划线单词转驼峰大写单词
func UnderscoreToUpperCamelCase(s string) string {
	//需要先把所有的下划线替换了
	s = strings.Replace(s, "_", " ", -1)      //把所有的都替换成空格
	s = cases.Title(language.Tag{}).String(s) //会将单词转为首字母大写
	return strings.ReplaceAll(s, " ", "")
}

// UnderscoreToLowerCameCase 下划线单词转小写驼峰单词
func UnderscoreToLowerCameCase(s string) string {
	s = UnderscoreToUpperCamelCase(s) //拿到所有首字母大写单词
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// CamelCaseUnderscore 驼峰转下划线
func CamelCaseUnderscore(s string) string {
	var result = make([]rune, 0, len(s))
	for i, value := range s {
		if i == 0 {
			result = append(result, unicode.ToLower(value))
			continue
		}
		if unicode.IsUpper(value) {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(value))
	}
	return string(result)
}
