package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	select {
	case v := <-ch:
		fmt.Println("Got:", v)
	default:
		fmt.Println("Channel is empty")
	}
}
