package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{Timeout: 5 * time.Second}

	req, err := http.NewRequest("GET", "https://httpbin.org/headers", nil)
	if err != nil {
		log.Fatalf("bad request: %v", err)
	}
	req.Header.Set("X-Student", "Zuhro")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("unexpected status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("read failed: %v", err)
	}
	fmt.Println(string(body))

	var result struct {
		Headers map[string]string `json:"headers"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("decode failed: %v", err)
	}

	if v, ok := result.Headers["X-Student"]; ok {
		fmt.Println("OK: header X-Student is present, value =", v)
	} else {
		fmt.Println("FAIL: header X-Student not found (such a shame!)")
	}
}
