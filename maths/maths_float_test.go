package maths

import (
	"fmt"
	"testing"
)

func TestRound(t *testing.T) {
	expected := 3.1455

	v := Round(3.14545, 4)

	if expected != v {
		t.Errorf("expected %f, got %f\n", expected, v)
	}
	fmt.Println(v)
}
