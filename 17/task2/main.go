package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	u := User{
		Name:  "Voldemort",
		Age:   71,
		Email: "ali@mail.tj",
	}

	data, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(data))
}
