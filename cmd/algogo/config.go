package main

import (
	"flag"
	"fmt"

	"github.com/Shogun-Sun/AlgoGo/pkg/sorting"
)

type Config struct {
	Algorithm string
	Size      int
	Verbose   bool
	Order     sorting.SortOrder
}

func ParseFlags() (Config, error) {
	var cfg Config
	var orderStr string

	flag.StringVar(&cfg.Algorithm, "algo", "stupid", "Алгоритм сортировки(stupid, bubble)")
	flag.IntVar(&cfg.Size, "size", 100, "Размер массива")
	flag.BoolVar(&cfg.Verbose, "verbose", false, "Вывод массива до и после сортировки")
	flag.StringVar(&orderStr, "order", "asc", "Порядок сортировки: asc (по возрастанию) или desc (по убыванию)")

	flag.Parse()

	switch orderStr {
	case "asc":
		cfg.Order = sorting.Ascending
	case "desc:":
		cfg.Order = sorting.Descending
	default:
		return Config{}, fmt.Errorf("неизвестный порядок сортировки '%s': используйте 'asc' или 'desc'", orderStr)
	}

	return cfg, nil
}
