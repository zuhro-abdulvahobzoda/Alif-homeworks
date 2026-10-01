package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	var wg sync.WaitGroup
	count := 100_000

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Millisecond)
		}()
	}

	wg.Wait()
	duration := time.Since(start)

	fmt.Printf("Time taken to spawn and wait for %d goroutines: %v\n", count, duration)
}
