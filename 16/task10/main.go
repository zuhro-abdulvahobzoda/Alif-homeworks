package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	results := make(chan int)
	var wg sync.WaitGroup

	numbers := []int{1, 2, 3}

	for _, n := range numbers {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			time.Sleep(time.Duration(val*100) * time.Millisecond)
			results <- val
		}(n)
	}

	// Helper goroutine to close the channel when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Read and sum up all values from results
	sum := 0
	for val := range results {
		sum += val
	}

	fmt.Printf("Total Sum: %d\n", sum)
}
