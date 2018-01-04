package sort_test

import (
	"testing"

	"github.com/bethanyg/golearners/sort"
)

func TestMergeSort(t *testing.T) {

	testCases := []struct {
		name string
		list []int
	}{
		{
			"nil slice",
			nil,
		},
		{
			"empty slice",
			[]int{},
		},
		{
			"single integer",
			[]int{100},
		},
		{
			"already sorted list",
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			"positive integers in reverse order",
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			"random initial list",
			[]int{4, 6, 0, 8, 2, 9, 1, 5, 3, 7},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := sort.MergeSort(tc.list)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if isSorted(got) == false {
				t.Errorf("slice not in numeric order: %v", got)
			}
		})
	}

}

// IsSorted returns true if all integers in a slice are sorted in increasing numeric order.
// Nil, empty or single element slices are classified as sorted.
func isSorted(a []int) bool {
	if a == nil {
		return true
	}

	if len(a) <= 1 {
		return true
	}

	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			return false
		}
	}

	return true
}
