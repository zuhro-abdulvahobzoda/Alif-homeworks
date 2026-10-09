package main

import (
	"fmt"
	"net/http"
	"time"
)

func classify(status int) string {
	switch {
	case status >= 200 && status < 300:
		return "SUCCESS"
	case status >= 400 && status < 500:
		return "Client error"
	case status >= 500 && status < 600:
		return "Server error"
	default:
		return "smth else idk"
	}
}

func main() {
	client := &http.Client{Timeout: 5 * time.Second}

	codes := []int{200, 201, 404, 500, 503}

	for _, code := range codes {
		url := fmt.Sprintf("https://httpbin.org/status/%d", code)

		resp, err := client.Get(url)
		if err != nil {
			fmt.Printf("%d: request failed: %v\n", code, err)
			continue
		}
		resp.Body.Close()

		fmt.Printf("%d: %s\n", resp.StatusCode, classify(resp.StatusCode))
	}
}
