package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Company struct {
	Name string `json:"name"`
}

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Company Company `json:"company"`
}

func main() {
	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("unexpected status: %s", resp.Status)
	}

	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		log.Fatalf("decode failed: %v", err)
	}

	for _, u := range users {
		fmt.Printf("%s - %s\n", u.Name, u.Company.Name)
	}
}
