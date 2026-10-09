package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type NewPost struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

type CreatedPost struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

func main() {
	client := &http.Client{Timeout: 5 * time.Second}

	payload := NewPost{
		Title:  "My first post",
		Body:   "Hello from Go!",
		UserID: 1,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("marshal failed: %v", err)
	}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("bad request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	var created CreatedPost
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		log.Fatalf("decode failed: %v", err)
	}

	fmt.Println("Status code:", resp.StatusCode)
	fmt.Println("Created post id:", created.ID)
}
