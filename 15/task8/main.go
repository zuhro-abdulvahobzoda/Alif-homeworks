package main

import (
	"fmt"
)

func generateNumbers(n int, ch chan int) {
	for i := 1; i <= n; i++ {
		ch <- i
	}

	close(ch)
}

func main() {
	ch := make(chan int)

	go generateNumbers(12, ch)

	for i := range ch {
		fmt.Println(i)
	}
}
