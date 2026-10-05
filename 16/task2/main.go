package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan string)
	ch2 := make(chan string)

	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		ch1 <- "from first channel"
	}()

	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		ch2 <- "from second channel"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}

	wg.Wait()

}
