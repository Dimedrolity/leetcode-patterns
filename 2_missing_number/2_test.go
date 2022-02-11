package missing_number

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	input := []int{3, 0, 1}
	want := 2
	got := missingNumber(input)
	if got != want {
		t.Fatalf(fmt.Sprint("got ", got))
	}
}
