package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	client := &http.Client{Timeout: 5 * time.Second}

	params := url.Values{}
	params.Set("userId", "3")

	endpoint := "https://jsonplaceholder.typicode.com/posts?" + params.Encode()

	resp, err := client.Get(endpoint)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("unexpected status: %s", resp.Status)
	}

	var posts []Post
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		log.Fatalf("decode failed: %v", err)
	}

	fmt.Println("Posts by user 3:", len(posts))
	for i, p := range posts {
		fmt.Printf("%d. %s\n", i+1, p.Title)
	}
}
