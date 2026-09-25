package sorting

import "math/rand"

type SortOrder string

const (
	Ascending  SortOrder = "asc"
	Descending SortOrder = "desc"
)

func GenerateSlice(size int, bounds ...int) []int {
	min, max := 0, 100

	if len(bounds) == 1 {
		max = bounds[0]
	} else if len(bounds) >= 2 {
		min = bounds[0]
		max = bounds[1]
	}

	slice := make([]int, size)

	for i := range size {
		slice[i] = min + rand.Intn(max-min+1)
	}
	return slice
}
