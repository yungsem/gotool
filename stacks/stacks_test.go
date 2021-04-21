package stacks

import (
	"fmt"
	"runtime/debug"
	"testing"
)

func TestStack(t *testing.T) {
	debug.PrintStack()
	sl := Stack(0)
	fmt.Println(sl)
}


func BenchmarkFormat(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 要测试的操作
		Stack(0)
	}
}