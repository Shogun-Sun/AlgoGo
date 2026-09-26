package test

import (
	"testing"

	"github.com/Shogun-Sun/AlgoGo/pkg/sorting"
)

func BenchmarkStupidSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := sorting.GenerateSlice(10000)
		b.StartTimer()

		sorting.Sort(data, sorting.StupidAlgorithm, sorting.Ascending)
	}
}
