package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// Print active goroutines before starting
	fmt.Println("Before launch:", runtime.NumGoroutine())

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
		}()
	}

	// Print active goroutines right after launching
	fmt.Println("After launching 50 goroutines:", runtime.NumGoroutine())

	wg.Wait()

	// Print active goroutines after all completed
	fmt.Println("After wg.Wait():", runtime.NumGoroutine())
}
