package main

import (
	"os"
	"runtime/trace"

	"github.com/bethanyg/golearners/sort"
)

// Data is an arbitrary set of lists that we can use to test our multi-sort
// approaches.
var data = [][]int{
	[]int{77, 26, 84, 39, 9, 52, 64, 41, 18, 86, 3, 12, 21, 41, 66, 63, 78, 67, 13, 73, 81, 55, 70, 0, 62, 74, 51, 8, 60, 17, 40, 96, 74, 64, 81, 38, 7, 22, 51, 10, 17, 31, 54, 46, 10},
	[]int{42, 99, 34, 48, 34, 76, 87, 76, 36, 16, 59, 67, 69, 14, 77, 36, 37, 36, 71, 64, 88, 53, 31, 55, 0, 20, 6, 27, 2, 54, 78, 40, 76, 85, 50, 15, 88, 29, 67, 22, 62},
	[]int{35, 11, 82, 81, 79, 30, 10, 82, 38, 69, 30, 76, 65, 39, 69, 33, 77, 4, 44, 1, 25, 97, 38, 71, 38, 35, 29, 4, 49, 61, 35, 28, 96, 74, 20, 32, 61, 0, 19, 69, 15, 54, 38},
	[]int{73, 11, 48, 10, 63, 92, 4, 92, 97, 11, 67, 76, 68, 67, 43, 42, 91, 15, 24, 67, 39, 38, 13, 14, 54, 82, 22, 19, 33, 62, 63, 2, 76},
	[]int{50, 85, 81, 57, 33, 58, 93, 78, 65, 10, 38, 15, 45, 22, 36, 44, 1, 79, 8, 42, 4, 91, 46, 30, 20, 84, 83, 18, 53, 4, 25, 75, 45, 68, 11, 78, 53, 79, 94, 97, 18, 9, 31, 39, 61, 88, 60, 58, 79, 61, 56, 88, 51},
	[]int{49, 69, 76, 15, 20, 79},
}

// Packet is a type that holds both an index and a list. It can be used both
// to submit lists for sorting and for returning the sorted list.
type packet struct {
	num  int
	list []int
}

// MergeSortMulti takes a slice of lists and sorts each list within the slice.
func MergeSortMulti(lists [][]int) [][]int {

	resChan := make(chan packet, len(lists))

	// Loop over the lists and create a go-routine to sort each of them. Each go-routine
	// writes the sorted list back to a channel.
	for idx := range lists {
		go func(li int, rc chan packet) {
			sorted := sort.MergeSort(lists[li])
			rc <- packet{num: li, list: sorted}
		}(idx, resChan)
	}

	// We know how many results we need to read from the results channel. Once we have
	// all the results we can close the channel.
	res := make([][]int, len(lists))
	for idx := 0; idx < len(lists); idx++ {
		p := <-resChan
		res[p.num] = p.list
	}
	close(resChan)

	// We can now return our sorted results.
	return res
}

func main() {

	// start tracing
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	// sort all the things
	MergeSortMulti(data)
}
