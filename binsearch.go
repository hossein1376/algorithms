package algorithms

import (
	"fmt"
)

func binSearch(slice []int, target int) int {
	l := len(slice)
	if l < 2 {
		panic(fmt.Errorf(
			"slice has %d elements, expected at least 2", l),
		)
	}

	low := 0
	high := l - 1

	for low <= high {
		mid := (low + high) / 2
		guess := slice[mid]

		if target == guess {
			return mid
		} else if target < guess {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
