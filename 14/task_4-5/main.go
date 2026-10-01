package main

import (
	"fmt"
	"sync"
	"time"
)

func printSlowly(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Set, go: ", name)
	time.Sleep(1 * time.Second)
	fmt.Println("Ready: ", name)
}

func greet(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Hello, %s!\n", name)
	time.Sleep(50 * time.Millisecond)

}

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	// these names are taken from Sonic the hedgehog franchise!
	names := []string{"Sonic", "Tails", "Knuckles", "Amy", "Dr. Eggman", "Dr.Starline", "Surge", "Whisper", "Tangle"}

	for _, name := range names {
		wg.Add(1)
		go printSlowly(name, &wg)
	}

	for _, name := range names {
		wg.Add(1)
		go greet(name, &wg)
		time.Sleep(300 * time.Millisecond)
	}

	wg.Wait()
	fmt.Println("Duration of the task: ", time.Since(start))
}
