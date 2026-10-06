package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func main() {
	var b Book
	input := `{"title":"Go для начинающих","author":"Иван Петров","year":2023}`

	err := json.Unmarshal([]byte(input), &b)
	if err != nil {
		log.Println("An error occured: %w", err)
		return
	}

	fmt.Println(b.Title)
	fmt.Println(b.Author)
	fmt.Println(b.Year)
}
