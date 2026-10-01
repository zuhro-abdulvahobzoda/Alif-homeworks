package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	go func() {
		ch <- "HI from goroutine!" // Task 1
	}()

	msg := <-ch
	fmt.Println(msg)

}
