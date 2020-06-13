package strings

import "strings"

func LowerCaseFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	first := s[0:1]
	rest := s[1:]
	return strings.ToLower(first) + rest
}
