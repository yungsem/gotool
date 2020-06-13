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
