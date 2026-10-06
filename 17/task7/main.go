package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	// no error would occur as if the city field in json doesn't exist at all
}

func main() {
	var p Person
	input := `{"name":"Алия","age":20,"city":"Душанбе"}`

	err := json.Unmarshal([]byte(input), &p)
	if err != nil {
		log.Println("An error occured: %w", err)
		return
	}

	fmt.Println(p.Name)
	fmt.Println(p.Age)

}
