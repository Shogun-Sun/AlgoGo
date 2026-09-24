package main

import (
	"fmt"

	"github.com/Shogun-Sun/AlgoGo/pkg/sorting"
)

func main() {
	fmt.Println(sorting.BubbleSort([]int{1, 2, 3, 4, 5, 6, 7, 8}, sorting.Ascending))
	fmt.Println(sorting.StupidSort([]int{1, 2, 3, 4, 5, 6, 7, 8}, sorting.Ascending))
}
