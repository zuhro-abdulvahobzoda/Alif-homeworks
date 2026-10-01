package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)

	v1, ok1 := <-ch
	fmt.Println(v1, ok1)
	v2, ok2 := <-ch
	fmt.Println(v2, ok2)
	v3, ok3 := <-ch
	fmt.Println(v3, ok3)
	_, ok4 := <-ch // false because the channel is closed and EMPTY
	fmt.Println(ok4)

}
