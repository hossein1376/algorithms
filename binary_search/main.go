package main

import (
	"fmt"
)

func main() {
	list := []int{1, 2, 3, 4, 5}
	fmt.Println(binSearch(list, 3))
}

func binSearch(list []int, i int) int {
	low, high := 0, len(list)-1

	for low <= high {
		mid := (low + high) / 2

		if list[mid] == i {
			return mid
		}
		if list[mid] < i {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
