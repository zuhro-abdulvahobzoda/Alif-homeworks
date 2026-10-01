package main

import (
	"fmt"
	"time"
)

func printSlowly(name string) {
	fmt.Println("Set, go: ", name)
	time.Sleep(1 * time.Second)
	fmt.Println("Ready: ", name)
}

func main() {
	start := time.Now()
	// printSlowly("Angela")
	// printSlowly("Ben")				task 1
	// printSlowly("Tom")

	go printSlowly("Angela")
	go printSlowly("Ben") //task 2
	go printSlowly("Tom")

	/* Ответ: горутины попросту не смогут спевать вывести
	что-то на терминал потому что что main заканчивается
	раньше.
	*/

	time.Sleep(1 * time.Second) // task 3
	fmt.Println("Duration of a task: ", time.Since(start))
}
