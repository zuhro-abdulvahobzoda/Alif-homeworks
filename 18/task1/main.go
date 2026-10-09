package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatal("Request error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Content-Type:", resp.Header.Get("Content-Type"))
	fmt.Println("Code status:", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Read error:", err)
		return
	}

	fmt.Println("Content (body):", string(body))

}
