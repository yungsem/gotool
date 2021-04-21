package errorf

import (
	"errors"
	"fmt"
	"testing"
)

func TestFormat(t *testing.T) {
	err := errors.New("test error")
	fmt.Println(Format(err))
}

func BenchmarkFormat(b *testing.B) {
	b.ResetTimer()
	err := errors.New("test error")
	for i := 0; i < b.N; i++ {
		// 要测试的操作
		Format(err)
	}
}