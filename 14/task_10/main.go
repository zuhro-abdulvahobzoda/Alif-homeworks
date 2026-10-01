package main

import (
	"fmt"
	"sync"
)

func orderOfExecution(n int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= n; i++ {
		fmt.Print(i, " ")
	}
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 3; i++ {
		wg.Add(1)
		go orderOfExecution(10, &wg)
		i++
	}
	wg.Wait()
	fmt.Println("Done bru")
}
