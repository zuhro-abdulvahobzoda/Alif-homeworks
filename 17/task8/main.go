package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Item struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Count int     `json:"count"`
}

func main() {
	var l Item
	input := `{"name":"Ноутбук","price":4500.5}`

	err := json.Unmarshal([]byte(input), &l)
	if err != nil {
		log.Println("An error occured: %w", err)
		return
	}

	fmt.Println(l.Name)
	fmt.Println(l.Price)
	fmt.Println(l.Count) // the Count field of Item structure will take the default value because json does not have the value

}
