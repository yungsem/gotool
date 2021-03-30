package strings

import (
	"regexp"
	"strings"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

// ToSnakeCaseLower 将 str 转换成小写的蛇形格式
// 如：realName -> real_name
// 如：RealName -> real_name
func ToSnakeCaseLower(str string) string {
	return strings.ToLower(toSnakeCase(str))
}

// ToSnakeCaseLower 将 str 转换成大写的蛇形格式
// 如：realName -> REAL_NAME
// 如：RealName -> REAL_NAME
func ToSnakeCaseUpper(str string) string {
	return strings.ToUpper(toSnakeCase(str))
}

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return snake
}

// LowerCaseFirstLetter 将 s 的首字母变成小写
// 如果 RealName -> realName
func LowerCaseFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	first := s[0:1]
	rest := s[1:]
	return strings.ToLower(first) + rest
}

