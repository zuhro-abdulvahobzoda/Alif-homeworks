package main

import "fmt"

func tryRead(ch chan string) {
	select {
	case msg := <-ch:
		fmt.Printf("Read message: %s\n", msg)
	default:
		fmt.Println("Empty")
	}
}

func main() {
	ch := make(chan string, 1)
	ch <- "Hello"

	tryRead(ch)
	tryRead(ch)
	tryRead(ch)
}
