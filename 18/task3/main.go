package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Todo struct {
	Title string `json:"title"`
}

func fetchTodo(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("An error occured during request %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("URL: %s\nCode status: %d\n", url, resp.StatusCode)

	switch resp.StatusCode {
	case http.StatusOK: // 200
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Decoding error: %v\n", err)
			return
		}

		var todo Todo
		if err := json.Unmarshal(body, &todo); err != nil {
			log.Printf("JSON parsing error: %v\n", err)
			return
		}

		fmt.Printf("Title: %s\n\n", todo.Title)

	case http.StatusNotFound: // 404
		fmt.Println("Task not found\n")

	default:
		fmt.Printf("Not EXPECTED status code: %d\n\n", resp.StatusCode)
	}
}

func main() {
	urls := []string{
		"https://jsonplaceholder.typicode.com/todos/5",
		"https://jsonplaceholder.typicode.com/todos/99999",
	}

	for _, url := range urls {
		fetchTodo(url)
	}
}
