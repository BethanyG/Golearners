package sort

import (
	"sync"
)

// MergeSortMulti takes a slice of lists and sorts each list within the slice.
func MergeSortMulti(lists [][]int) [][]int {

	res := make([][]int, len(lists))

	wg := sync.WaitGroup{}

	for idx := range lists {
		wg.Add(1)
		go func(li int) {
			res[li] = MergeSort(lists[li])
			wg.Done()
		}(idx)
	}

	wg.Wait()
	return res
}
