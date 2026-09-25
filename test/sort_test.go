package test

import (
	"testing"

	"github.com/Shogun-Sun/AlgoGo/pkg/sorting"
)

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := sorting.GenerateSlice(10000)
		b.StartTimer()

		sorting.BubbleSort(data, sorting.Descending)
	}
}

func BenchmarkStupidSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := sorting.GenerateSlice(10000)
		b.StartTimer()

		sorting.StupidSort(data, sorting.Descending)
	}
}
