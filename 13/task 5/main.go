package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func writeCSV(filename string, products []Product) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	w.Write([]string{"Name", "Price", "Quantity"})
	for _, p := range products {
		if err := w.Write([]string{p.Name, strconv.FormatFloat(p.Price, 'f', -1, 64), strconv.Itoa(p.Quantity)}); err != nil {
			return err
		}
	}
	return nil

}

func readCSV(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return csv.NewReader(f).ReadAll()
}

func main() {
	products := []Product{
		{"Apple", 9.5, 7},
		{"Pear", 9.5, 9},
		{"Ananas", 25.2, 4},
		{"Lime", 5.1, 17},
		{"Orange", 14.0, 23},
	}

	if err := writeCSV("products.csv", products); err != nil {
		fmt.Println("Ошибка записи:", err)
		return
	}

	datas, err := readCSV("products.csv")
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}

	for i, rec := range datas {
		fmt.Printf("%-10s %-15s %s\n", rec[0], rec[1], rec[2])
		if i == 0 {
			fmt.Println("-----------------------------------")
		}
	}
}
