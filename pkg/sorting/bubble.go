package sorting

func BubbleSort(elements []int, sortOrder string) []int {
	for i := 0; i < len(elements); i++ {
		for j := 0; j < len(elements)-i-1; j++ {
			if sortOrder == "asc" {
				if elements[j] > elements[j+1] {
					elements[j], elements[j+1] = elements[j+1], elements[j]
				}
			} else {
				if elements[j] < elements[j+1] {
					elements[j], elements[j+1] = elements[j+1], elements[j]
				}
			}
		}
	}
	return elements
}
