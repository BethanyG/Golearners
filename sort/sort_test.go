package sort_test

import (
	"testing"

	"github.com/bethanyg/golearners/sort"
)

func TestMergeSort(t *testing.T) {

	testCases := []struct {
		name string
		list []int
		want []int
	}{
		{
			"positive integers in reverse order",
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			"already sorted list",
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			"random initial list",
			[]int{4, 6, 0, 8, 2, 9, 1, 5, 3, 7},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := sort.MergeSort(tc.list)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if isEqual(got, tc.want) == false {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}

}

// IsEqual is a helper function that compares two slices of integers and returns
// true if they are equal.
func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
