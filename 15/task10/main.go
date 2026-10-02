package main

import (
	"fmt"
)

func sendNames(names []string, ch chan<- string) {
	go func() {
		for _, name := range names {
			ch <- name
		}
		close(ch)
	}()
}

func main() {
	students := []string{"Аня", "Борис", "Вика", "Данил", "Ева"}
	ch := make(chan string)

	sendNames(students, ch)

	var result []string
	for name := range ch {
		result = append(result, name)
	}

	// Output total count and the resulting slice
	fmt.Printf("Total names received: %d\n", len(result))
	fmt.Printf("Result slice: %v\n", result)
}
