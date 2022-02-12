package find_disappeared_numbers

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func Test(t *testing.T) {
	input := []int{4, 3, 2, 7, 8, 2, 3, 1}
	want := []int{5, 6}

	got := findDisappearedNumbers(input)
	sort.Ints(got)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf(fmt.Sprint("got ", got))
	}
}
