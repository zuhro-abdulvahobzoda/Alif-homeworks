package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3
	// ch <- 4 выдает ошибку т.к мы уже задали обьём канала (или буферизовали??) и больше заданного обьёма превысить не можем

	fmt.Printf("Length of this channel: %d || capacity: %d\n", len(ch), cap(ch))

	val1 := <-ch
	val2 := <-ch
	val3 := <-ch

	fmt.Printf("%d, %d, %d\n", val1, val2, val3)
	fmt.Printf("Current length of this channel: %d || capacity: %d\n", len(ch), cap(ch))
}
