package maths

import (
	"fmt"
	"testing"
)

func TestRound(t *testing.T) {
	expected := 3.1455

	f := float64(9)/float64(10)
	v := Round(f, 2)

	if expected != v {
		t.Errorf("expected %f, got %f\n", expected, v)
	}
	fmt.Println(v)
}
