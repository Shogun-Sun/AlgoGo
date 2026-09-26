package main

import (
	"flag"
	"fmt"

	"github.com/Shogun-Sun/AlgoGo/pkg/sorting"
)

type Config struct {
	Algorithm sorting.AlgorithmType
	Size      int
	Verbose   bool
	Order     sorting.SortOrder
	Min       int
	Max       int
}

func ParseFlags() (Config, error) {
	var cfg Config
	var orderStr string
	var algoStr string

	flag.StringVar(&algoStr, "algo", "stupid", "Алгоритм сортировки(stupid, bubble)")
	flag.IntVar(&cfg.Size, "size", 100, "Размер массива")
	flag.BoolVar(&cfg.Verbose, "verbose", false, "Вывод массива до и после сортировки")
	flag.StringVar(&orderStr, "order", "asc", "Порядок сортировки: asc (по возрастанию) или desc (по убыванию)")
	flag.IntVar(&cfg.Min, "min", 0, "Нижняя граница генерации (включительно)")
	flag.IntVar(&cfg.Max, "max", 10, "Верхняя граница генерации (включительно)")

	flag.Parse()

	switch orderStr {
	case "asc":
		cfg.Order = sorting.Ascending
	case "desc":
		cfg.Order = sorting.Descending
	default:
		return Config{}, fmt.Errorf("неизвестный порядок сортировки '%s': используйте 'asc' или 'desc'", orderStr)
	}

	switch algoStr {
	case "stupid":
		cfg.Algorithm = sorting.StupidAlgorithm
	case "bubble":
		cfg.Algorithm = sorting.BubbleAlgorithm
	case "bubble-optimization":
		cfg.Algorithm = sorting.BubbleOptimizationAlgorithm
	default:
		return Config{}, fmt.Errorf("неизвестный алгоритм сортировки '%s'", algoStr)
	}

	return cfg, nil
}
