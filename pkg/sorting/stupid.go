package sorting

func StupidSort(elements []int, order SortOrder) []int {
	isAsc := order == Ascending
	i := 0

	for i < len(elements) {
		if i == 0 {
			i++
			continue
		}

		shouldSwap := (isAsc && elements[i-1] > elements[i]) ||
			(!isAsc && elements[i-1] < elements[i])

		if shouldSwap {
			elements[i], elements[i-1] = elements[i-1], elements[i]
			i--
		} else {
			i++
		}
	}

	return elements
}
