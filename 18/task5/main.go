package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts")
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

	if len(posts) == 0 {
		fmt.Println("No posts returned")
		return
	}

	fmt.Println("Total posts:", len(posts))
	fmt.Println("First title:", posts[0].Title)
	fmt.Println("Last title: ", posts[len(posts)-1].Title)
}
