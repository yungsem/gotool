package strings

import (
	"fmt"
	"testing"
)

func TestLowerCaseFirstLetter(t *testing.T) {
	fmt.Println(LowerCaseFirstLetter("Hello"))
	fmt.Println(LowerCaseFirstLetter("hello"))
	fmt.Println(LowerCaseFirstLetter(""))
	fmt.Println(LowerCaseFirstLetter(" Hello"))
	fmt.Println(LowerCaseFirstLetter(" hello"))
	fmt.Println(LowerCaseFirstLetter("中间"))
}

func TestMD5(t *testing.T) {
	fmt.Println(MD5("hello"))
	fmt.Println(MD5("https://github.com/yungsem/gotool"))
	fmt.Println(MD5("https://stackoverflow.com/questions/47293975/should-i-error-check-close-on-a-response-body"))
}