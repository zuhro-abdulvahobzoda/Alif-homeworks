package main

func main() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	close(ch) // panic: close of closed channel

	ch <- 4 // send on closed channel

}
