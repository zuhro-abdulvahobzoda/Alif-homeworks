package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student struct {
	Name    string `json:"name"`
	Grade   int    `json:"grade"`
	Comment string `json:"cooment,omitempty"` // meaning if it's empty, json will ignore it
}

func main() {
	students := []Student{
		{"Alan", 89, "Excellent"},
		{"Nael", 67, ""},
	}

	data, err := json.MarshalIndent(students, "", "  ")
	if err != nil {
		log.Println("An error occurred: ", err)
	}
	fmt.Println(string(data))
}
