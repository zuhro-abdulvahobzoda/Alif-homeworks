package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	notAdded := 0
	added := 0
	for i := 1; i <= 5; i++ {

		select {
		case ch <- i:
			added++
		default:
			notAdded++
		}

	}

	fmt.Printf("Added: %d || Skipped: %d", added, notAdded)
}
