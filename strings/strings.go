package strings

import (
	"regexp"
	"strings"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func ToSnakeCaseLower(str string) string {
	return strings.ToLower(toSnakeCase(str))
}

func ToSnakeCaseUpper(str string) string {
	return strings.ToUpper(toSnakeCase(str))
}

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return snake
}

func LowerCaseFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	first := s[0:1]
	rest := s[1:]
	return strings.ToLower(first) + rest
}
