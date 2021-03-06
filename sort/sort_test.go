package sort_test

import (
	"fmt"
	"math/rand"
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
		{
			"random slice of 256 digits",
			rand.Perm(256),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := sort.MergeSort(tc.list)

			if isSorted(got) == false {
				t.Errorf("slice not in numeric order: %v", got)
			}
		})
	}

}

func BenchmarkMergeSort(b *testing.B) {
	benchmarks := []int{10, 100, 1000, 10000, 100000, 1000000}

	for _, bm := range benchmarks {
		b.Run(fmt.Sprintf("Benchmark %d", bm), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sort.MergeSort(rand.Perm(bm))
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
