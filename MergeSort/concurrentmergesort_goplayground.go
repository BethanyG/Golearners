package main

import (
  "fmt"
  "sync"
)

var buffSorter = make(chan struct{}, 100)
var buffListSorter = make(chan []int, 100)


//func ListSorter takes a list of lists and sorts them,
//returning the list in it's original order.
func ListSorter(lists [][]int) [][]int {
  lwg := sync.WaitGroup{}
  completedList := make([][]int, len(lists))
  lwg.Add(len(lists))

  for index, item := range lists {
        go func(index int, item []int) {
            sorted := MergeSortMulti(item)
            fmt.Println("STARTING:: ", index)
            //buffListSorter <- sorted
	    completedList[index] = sorted
	    fmt.Println("WRITING:: ", index)
            lwg.Done()
        }(index, item)
  }
  lwg.Wait()
  return completedList
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


func main() {
    s := [][]int{[]int{77, 26, 84, 39, 9, 52, 64, 41, 18, 86, 3, 12, 21, 41, 66, 63, 78, 67, 13, 73, 81, 55, 70, 0, 62, 74, 51, 8, 60, 17, 40, 96, 74, 64, 81, 38, 7, 22, 51, 10, 17, 31, 54, 46, 10},
[]int{42, 99, 34, 48, 34, 76, 87, 76, 36, 16, 59, 67, 69, 14, 77, 36, 37, 36, 71, 64, 88, 53, 31, 55, 0, 20, 6, 27, 2, 54, 78, 40, 76, 85, 50, 15, 88, 29, 67, 22, 62},
[]int{35, 11, 82, 81, 79, 30, 10, 82, 38, 69, 30, 76, 65, 39, 69, 33, 77, 4, 44, 1, 25, 97, 38, 71, 38, 35, 29, 4, 49, 61, 35, 28, 96, 74, 20, 32, 61, 0, 19, 69, 15, 54, 38},
[]int{73, 11, 48, 10, 63, 92, 4, 92, 97, 11, 67, 76, 68, 67, 43, 42, 91, 15, 24, 67, 39, 38, 13, 14, 54, 82, 22, 19, 33, 62, 63, 2, 76},
[]int{50, 85, 81, 57, 33, 58, 93, 78, 65, 10, 38, 15, 45, 22, 36, 44, 1, 79, 8, 42, 4, 91, 46, 30, 20, 84, 83, 18, 53, 4, 25, 75, 45, 68, 11, 78, 53, 79, 94, 97, 18, 9, 31, 39, 61, 88, 60, 58, 79, 61, 56, 88, 51},
[]int{49, 69, 76, 15, 20, 79},
[]int{34, 6, 53, 98, 80, 13, 43, 13, 46, 87, 91, 23, 76, 18, 78, 36, 44, 52, 97, 46, 53, 91, 23, 42, 63, 46, 90, 35, 2, 88, 92, 56, 6, 43, 20, 49, 40, 30},
[]int{1, 8, 31, 64, 80, 68, 11, 67, 67, 76, 12, 67, 30, 97, 78, 51, 25, 24, 23, 90},
[]int{24, 20, 87, 59, 12, 72, 83, 93, 71, 30, 11, 11, 56, 29, 55, 0, 36, 2, 75, 31, 84, 62, 65, 55, 86, 75, 28, 87, 75, 6, 53, 22, 48, 39, 67, 12, 16, 9, 14},
[]int{8, 77, 88, 3, 59, 83, 75, 67, 23, 53, 98, 83, 33, 52, 4, 13, 51, 0, 43, 17, 67, 41, 10, 78, 3, 8, 1, 22, 34, 18, 11},
[]int{28, 84, 72, 67, 70, 20, 53, 11, 95, 3, 83, 24, 3, 94, 98, 97, 91, 3, 65, 65, 0, 42, 21, 14, 61, 40, 17, 16, 93, 65, 86, 95, 98, 73, 73, 69, 53, 48, 98, 73, 13, 54, 9, 37, 18, 71, 28, 41, 93, 95, 18, 70, 96},
[]int{39, 57, 29, 83, 55, 90, 91},
[]int{75, 9, 11, 96, 34, 56, 10, 8, 83, 35, 73, 24, 71, 49, 33, 30, 36, 75, 35, 49, 94, 97, 93, 97, 74, 96, 85, 69, 51, 23, 63, 83, 6, 32, 54, 13, 19, 82, 64, 19, 97, 92, 75, 96, 88, 55},
[]int{36, 59, 59, 0, 92, 32, 23, 0, 14, 34, 25, 21, 91, 83, 86, 18, 76, 60, 76, 41},
[]int{17, 53, 96, 99, 82, 67, 92, 31, 26, 70, 26, 22, 18, 31, 22, 99, 46, 40, 92, 3, 46, 8, 73, 24, 74, 67, 90, 24, 8, 53, 79, 22, 60, 83, 6, 70, 22},
[]int{56, 24, 70, 62, 46, 94, 38, 52, 64, 86, 34, 30, 81, 55, 86, 30, 67, 4, 37, 33, 16, 68, 32, 79},
[]int{72, 21, 65, 6, 57, 84, 8, 82, 92, 69, 82, 77, 25, 35, 85, 1, 56, 93, 40, 89, 51, 86, 69, 94, 69, 18, 73, 1, 46, 42, 22, 21, 4, 40, 12, 5, 96, 96, 11, 94, 4},
[]int{92, 32, 36, 90, 39, 22},
[]int{48, 10, 68, 15, 78, 76, 22, 37, 10, 67, 40, 56, 95},
[]int{12, 29, 77, 17, 45, 59, 56, 10, 73, 23, 94, 14, 70, 39, 78, 10, 56, 85, 72, 3, 7, 13, 54, 43, 25, 21, 1, 67, 57, 69, 29, 63, 73, 30, 60, 7, 14, 83, 24}}

 results := ListSorter(s)

 for _,v := range results {
       fmt.Println(v)
   }
}
