package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	resp, err := client.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		log.Fatal("Error request bru")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode)
		return
	}

	var todos []Todo
	if err := json.NewDecoder(resp.Body).Decode(&todos); err != nil {
		fmt.Printf("Failed to parse JSON: %v\n", err)
		return
	}

	var completed, notCompleted int
	for _, todo := range todos {
		if todo.Completed {
			completed++
		} else {
			notCompleted++
		}
	}

	fmt.Printf("Completed (completed: true): %d\n", completed)

}
