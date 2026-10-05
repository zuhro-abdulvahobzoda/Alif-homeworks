package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)

	go func() {
		defer close(even)
		defer close(odd)

		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				even <- i
			} else {
				odd <- i
			}
		}
	}()

	evenOpen := true
	oddOpen := true

	for evenOpen || oddOpen {
		select {
		case val, ok := <-even:
			if !ok {
				evenOpen = false
			} else {
				fmt.Printf("even: %d\n", val)
			}
		case val, ok := <-odd:
			if !ok {
				oddOpen = false
			} else {
				fmt.Printf("odd: %d\n", val)
			}
		}
	}
}
