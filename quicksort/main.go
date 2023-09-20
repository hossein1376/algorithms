package main

import (
	"fmt"
)

func main() {
	list := []int{4, 2, 1, 5, 3}
	fmt.Println(quickSort(list))
}

func quickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}

	pivot := list[0]
	less, greater := []int{}, []int{}
	for _, num := range list[1:] {
		if pivot > num {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}

	less = append(quickSort(less), pivot)
	greater = append(quickSort(greater))

	return append(less, greater...)
}
