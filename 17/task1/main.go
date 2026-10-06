package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Book struct {
	Title  string
	Author string
	Year   int
	Price  float64
}

func main() {

	book := Book{"Harry Potter", "J.K.Rowling", 1997, 10.5}

	data, err := json.Marshal(book)
	if err != nil {
		log.Println("An error occured: %w", err)
		return
	}

	fmt.Println(string(data))
}
