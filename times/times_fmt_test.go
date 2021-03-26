package times

import (
	"fmt"
	"testing"
	"time"
)

func TestFormatDate(t *testing.T) {
	ti:= FormatDate(time.Now())
	fmt.Printf("%s\n", ti)
}

func TestFormatDateTime(t *testing.T) {
	ti:= FormatDateTime(time.Now())
	fmt.Printf("%s\n", ti)
}