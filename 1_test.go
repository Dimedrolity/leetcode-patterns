package leetcode_patterns

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	input := []int{1, 2, 3, 1}

	if got := containsDuplicate(input); got != true {
		t.Fatalf(fmt.Sprint("got ", got))
	}
}
