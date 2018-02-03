// Package sublist contains a function to determine whether or not one list is a sublist of another
package sublist

type Relation string

// IsSublist determines whether or not first list is a sublist of second list
func IsSublist(firstList []int, secondList []int) bool {

	lenFirstList := len(firstList)
	lenSecondList := len(secondList)

	if lenFirstList > lenSecondList {
		return false
	}

	for i := 0; i < lenSecondList-lenFirstList+1; i++ {
		sameValue := true
		var j int
		for j = 0; j < lenFirstList && sameValue; j++ {
			if firstList[j] != secondList[i+j] {
				sameValue = false
			}
		}
		if j == lenFirstList && sameValue {
			// is sublist
			return true
		}
	}

	return false
}

// Sublist tests whether first list is a sublist of second list
func Sublist(firstList []int, secondList []int) Relation {
	capacityFirstList := len(firstList)
	capacitySecondList := len(secondList)
	if capacityFirstList == 0 && capacitySecondList == 0 {
		return "equal"
	} else if capacityFirstList == 0 && capacitySecondList != 0 {
		return "sublist"
	} else if capacityFirstList != 0 && capacitySecondList == 0 {
		return "superlist"
	}

	isFirstsublist := IsSublist(firstList, secondList)
	isFirstSuperList := IsSublist(secondList, firstList)
	switch {
	case isFirstsublist && isFirstSuperList:
		return "equal"
	case isFirstsublist && !isFirstSuperList:
		return "sublist"
	case !isFirstsublist && isFirstSuperList:
		return "superlist"
	}

	return "unequal"
}
