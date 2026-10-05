package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	messages := make(chan string)

	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		timer := time.After(2 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				messages <- "ping"
			case <-timer:
				close(done)
				return
			}
		}
	}()

Loop:
	for {
		select {
		case msg := <-messages:
			fmt.Println(msg)
		case <-done:
			fmt.Println("End")
			break Loop
		}
	}
}
