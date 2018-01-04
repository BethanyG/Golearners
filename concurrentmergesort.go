//package ConcurrentMergeSort concurrently mergsorts a variable list of arrays.
package ConcurrentMergeSort

import (
  "fmt"
  "sync"
)

var buffSorter = make(chan struct{}, 100)


//func ListSorter takes a slice of slices and sorts each,
//returning the main slice in it's original order.
func ListSorter(lists [][]int) [][]int {
  lwg := sync.WaitGroup{}
  completedList := make([][]int, len(lists))
  lwg.Add(len(lists))

  for index, item := range lists {
        go func(index int, item []int) {
            sorted := MergeSortMulti(item)
            fmt.Println("STARTING:: ", index)
	    completedList[index] = sorted
	    fmt.Println("WRITING:: ", index)
            lwg.Done()
        }(index, item)
  }
  lwg.Wait()
  return completedList
}

  for item := range buffListSorter {
    completedList = append(completedList, item)
  }
  lwg.Wait()
  return completedList
  }
}

// MergeSortMulti sorts the slice items using Merge Sort Algorithm
func MergeSortMulti(items []int) []int {
    if len(items) <= 1 {
        return items
    }

    n := len(items) / 2

    wg := sync.WaitGroup{}
    wg.Add(2)

    var left, right []int

    select {
    case buffSorter <- struct{}{}:
        go func() {
            left = MergeSortMulti(items[:n])
            <-buffSorter
            wg.Done()
        }()
    default:
        left = MergeSortMulti(items[:n])
        wg.Done()
    }

    select {
    case buffSorter <- struct{}{}:
        go func() {
            right = MergeSortMulti(items[n:])
            <-buffSorter
            wg.Done()
        }()
    default:
        right = MergeSortMulti(items[n:])
        wg.Done()
    }

    wg.Wait()
    return merge(left, right)
}


func merge(left, right []int) []int {
	retVal := make([]int, 0, len(left)+len(right))

  for len(left) > 0 || len(right) > 0 {
    switch {
    case len(left) == 0 :
			return append(retVal, right...)
		case len(right) == 0 :
			return append(retVal, left...)
		case left[0] <= right[0] :
			retVal = append(retVal, left[0])
			left = left[1:]
		default:
			retVal = append(retVal, right[0])
			right = right[1:]
		}
	}
	return retVal
}
