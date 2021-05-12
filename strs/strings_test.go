package strs

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
}

func TestToSnakeCaseLower(t *testing.T) {
	fmt.Println(ToSnakeCaseLower("UserDetail"))
	fmt.Println(ToSnakeCaseLower("userDetail"))

	fmt.Println("---------------------------------")
	fmt.Println(ToSnakeCaseUpper("UserDetail"))
	fmt.Println(ToSnakeCaseUpper("userDetail"))
}