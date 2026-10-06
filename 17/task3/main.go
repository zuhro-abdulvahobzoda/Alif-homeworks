package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Product struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	InStock bool    `json:"in_stock"`
}

func main() {
	products := []Product{
		{000, "Apple", 10.5, true},
		{001, "Banana", 8.5, true},
		{002, "Orange", 12.0, false},
	}

	datas, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		log.Println("Error: ", err)
		return
	}

	fmt.Println(string(datas))
}
