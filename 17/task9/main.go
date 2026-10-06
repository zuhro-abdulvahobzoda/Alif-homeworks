package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Address struct {
	City   string `json:"city"`
	Street string `json:"street"`
}

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
}

func main() {
	// Go -> json

	p := Person{
		Name: "Junichi Kuroga",
		Age:  54,
		Address: Address{
			City:   "Tokyo",
			Street: "Takeshita Street",
		},
	}

	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Println("Error: ", err)
		return
	}

	fmt.Println(string(data))

	// json -> Go

	var person Person
	err2 := json.Unmarshal(data, &person)
	if err2 != nil {
		log.Println("An error occured: %w", err)
		return
	}

	fmt.Println("\nName: ", person.Name)
	fmt.Println("Age: ", person.Age)
	fmt.Println("City: ", person.Address.City)
	fmt.Println("Street: ", person.Address.Street)
}
