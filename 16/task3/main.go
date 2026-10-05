package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		// time.Sleep(3 * time.Second) == always will be timeout
		time.Sleep(1 * time.Second) // the channel has its value
		ch <- 42
	}()

	select {
	case msg := <-ch:
		fmt.Println("The channel has its value: ", msg)
		// case <-time.After(1 * time.Second):
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout")
	}

}
