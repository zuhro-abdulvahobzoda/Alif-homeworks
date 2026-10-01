package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int, 5)

	start1 := time.Now()
	for i := 1; i <= 5; i++ {
		ch1 <- i
	}
	duration1 := time.Since(start1)

	ch2 := make(chan int, 5)

	go func() {
		for i := 0; i < 6; i++ {
			time.Sleep(200 * time.Millisecond)
			<-ch2
		}
	}()

	start2 := time.Now()
	for i := 1; i <= 6; i++ {
		ch2 <- i
	}
	duration2 := time.Since(start2)

	fmt.Printf("Without goroutine (5 elements): %v\n", duration1)
	fmt.Printf("With goroutine 2 (6 elements): %v\n", duration2)
}
