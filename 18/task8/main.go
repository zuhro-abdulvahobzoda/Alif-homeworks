package main

import (
	"fmt"
	"net/http"
	"time"
)

const url = "https://httpbin.org/delay/2"

func main() {
	// Attempt 1: timeout (1s)
	shortClient := &http.Client{Timeout: 1 * time.Second}

	resp, err := shortClient.Get(url)
	if err != nil {
		fmt.Println("1s timeout, error:", err)
	} else {
		resp.Body.Close()
		fmt.Println("1s timeout, status:", resp.StatusCode)
	}

	// Attempt 2: timeout (5s)
	longClient := &http.Client{Timeout: 5 * time.Second}

	resp, err = longClient.Get(url)
	if err != nil {
		fmt.Println("5s timeout, error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("5s timeout, status:", resp.StatusCode)
}
