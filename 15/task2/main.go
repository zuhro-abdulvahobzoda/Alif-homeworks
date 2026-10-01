package main

import (
	"fmt"
)

func square(n int, ch chan int) {
	ch <- n * n
}

func main() {
	ch := make(chan int)

	go square(7, ch)

	result := <-ch
	fmt.Printf("7 in square = %d\n", result)
}
