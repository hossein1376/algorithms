package algorithms

func bubbleSort(slice []int) {
	l := len(slice)
	if l < 2 {
		return
	}

	back := 0
	front := 1
	sorted := false

	for !sorted {
		sorted = true
		for range slice[:l-1] {
			if slice[back] > slice[front] {
				swap(slice, back, front)
				sorted = false
			}
			back++
			front++
		}
		back = 0
		front = 1
	}

	return
}

func swap(s []int, a, b int) {
	temp := s[a]
	s[a] = s[b]
	s[b] = temp
}
