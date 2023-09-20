package main

import (
	"fmt"
)

func main() {
	list := []int{4, 2, 5, 1, 3}
	fmt.Println(selectionSort(list))
}

func selectionSort(list []int) []int {
	size := len(list)
	newArr := make([]int, size)

	for i := 0; i < size; i++ {
		smallest := findSmallest(list)
		newArr[i] = list[smallest]
		list = append(list[:smallest], list[smallest+1:]...)
	}
	return newArr
}

func findSmallest(list []int) int {
	smallest := 0
	for i := 0; i < len(list); i++ {
		if list[i] < list[smallest] {
			smallest = i
		}
	}
	return smallest
}
