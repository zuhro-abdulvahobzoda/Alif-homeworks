package catalog

import (
	_ "fmt"
)

type Product struct {
	Name  string
	Price float64
}

func TotalPrice(products []Product) float64 {
	total := 0.0
	for _, p := range products {
		total += p.Price
	}

	return total
}
