package main

import (
	"fmt"
	"log"

	"github.com/Shogun-Sun/AlgoGo/pkg/sorting"
)

func main() {
	cfg, err := ParseFlags()
	if err != nil {
		log.Fatalf("Ошибка конфигурации: %v", err)
	}

	arr := []int{5, 2, 9, 1, 3}
	fmt.Printf("Исходный массив: %v\n", arr)
	sorting.StupidSort(arr, cfg.Order)
	fmt.Printf("Отсортированный (%s): %v\n", cfg.Order, arr)
}
