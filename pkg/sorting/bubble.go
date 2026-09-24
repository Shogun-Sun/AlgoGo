package sorting

func BubbleSort(elements []int, order SortOrder) []int {
	isAsc := order == Ascending
	for i := range elements {
		for j := 0; j < len(elements)-i-1; j++ {
			shouldSwap := (isAsc && elements[j] > elements[j+1]) ||
				(!isAsc && elements[j] < elements[j+1])

			if shouldSwap {
				elements[j], elements[j+1] = elements[j+1], elements[j]
			}
		}
	}
	return elements
}
