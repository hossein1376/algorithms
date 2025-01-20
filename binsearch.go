package algorithms

func binSearch(slice []int, target int) int {
	l := len(slice)
	switch l {
	case 0:
		return -1
	case 1:
		if slice[0] == target {
			return 0
		}
		return -1
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
