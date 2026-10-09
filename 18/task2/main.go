package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func main() {
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	resp, err := client.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		log.Fatal("Error: ", err)
	}

	defer resp.Body.Close()

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		fmt.Println("Decoding error:", err)
		return
	}

	fmt.Printf("Name: %v", user.Name)
	fmt.Printf("Phone: %v", user.Phone)
	fmt.Printf("Email: %v", user.Email)
}
