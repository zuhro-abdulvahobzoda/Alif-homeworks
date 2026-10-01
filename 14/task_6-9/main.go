package main

import (
	"fmt"
	"sync"
	"time"
)

func countTo(n int, wg *sync.WaitGroup) { // task 6
	defer wg.Done()

	fmt.Printf("\nCounting from zero to %d...\n", n)
	for i := 0; i <= n; i++ {
		fmt.Println(i)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	// wg.Add(1)
	// go countTo(3, &wg)

	// wg.Add(1)
	// go countTo(5, &wg)

	// wg.Add(1)
	// go countTo(7, &wg)

	nums := []int{3, 5, 7} // task 7

	for _, num := range nums {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()

			fmt.Printf("\nCounting from zero to %d...\n", n)

			for i := 0; i <= n; i++ {
				fmt.Println(i)
				time.Sleep(10 * time.Millisecond)
			}
		}(num)
	}
	wg.Wait()
	fmt.Print("DONE!!\n")

	time.Sleep(1 * time.Second)

	fmt.Println("Just a normal anonimous function")

	for i := 1; i <= 8; i++ { // task 9
		func(id int) {
			fmt.Printf("Goroutine %d works\n", id)
			time.Sleep(time.Duration(id) * 50 * time.Millisecond)
			fmt.Printf("Gourutine %d finished\n", id)
		}(i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("\nAnonimous function with goroutine")

	for i := 1; i <= 8; i++ { // task 8
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d works\n", id)
			time.Sleep(time.Duration(id) * 50 * time.Millisecond)
			fmt.Printf("Gourutine %d finished\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Print("DONE!!")

}
