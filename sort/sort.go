// Package sort contains functions for sorting integer slices.
package sort

// MergeSort implements a basic merge sort. It takes a slice of integers
// and returns a new slice sorted in numeric order.
func MergeSort(list []int) []int {

	var res []int

	// Nil, empty, or short lists can be returned as is.
	if list == nil || len(list) <= 1 {
		return list
	}

	// If we have a list of two items return them in numeric order.
	if len(list) == 2 {
		if list[0] > list[1] {
			list[0], list[1] = list[1], list[0]
		}

		return list
	}

	// If our list contains more than two items split the lists.
	pp := len(list) / 2

	// Sort the two new lists and merge the results.
	a := MergeSort(list[:pp])
	b := MergeSort(list[pp:])
	res = merge(a, b)

	return res
}

// Merge is a helper function that merges two lists in increasing numerical order.
func merge(a, b []int) []int {

	if len(a) == 0 {
		return b
	}

	if len(b) == 0 {
		return a
	}

	pa := 0
	pb := 0
	c := make([]int, len(a)+len(b))

	for pc := 0; pa+pb < len(c); pc = pa + pb {

		if pa == len(a) {
			c[pc] = b[pb]
			pb++
			continue
		}

		if pb == len(b) {
			c[pc] = a[pa]
			pa++
			continue
		}

		if a[pa] < b[pb] {
			c[pc] = a[pa]
			pa++
		} else {
			c[pc] = b[pb]
			pb++
		}
	}

	return c
}
