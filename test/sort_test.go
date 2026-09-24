package test

import (
	"math/rand"
	"testing"

	"github.com/Shogun-Sun/AlgoGo/pkg/sorting"
)

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := generateSlice(10000)
		b.StartTimer()

		sorting.BubbleSort(data, "asc")
	}
}

func generateSlice(size int) []int {
	slice := make([]int, size)
	for i := range size {
		slice[i] = rand.Int()
	}
	return slice
}
