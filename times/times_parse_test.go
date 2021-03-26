package times

import (
	"fmt"
	"testing"
)

func TestParseDate(t *testing.T) {
	ti, err := ParseDate("2021-03-26")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%v\n", ti)
}

func TestParseDateTime(t *testing.T) {
	ti, err := ParseDateTime("2021-03-26 21:39:12")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%v\n", ti)
}