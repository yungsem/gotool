package strings

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func LowerCaseFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	first := s[0:1]
	rest := s[1:]
	return strings.ToLower(first) + rest
}

func MD5(s string) string {
	has := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", has)
}
