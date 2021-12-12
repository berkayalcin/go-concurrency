package main

import "fmt"

func main() {
	// Blocking read and write
	bufferedChan1 := make(chan int, 1)
	bufferedChan1 <- 1

	bufferedChan2 := make(chan int, 2)
	bufferedChan2 <- 2

	var done bool
	for !done {
		select {
		case c1Val := <-bufferedChan1:
			fmt.Println(c1Val)
		case c2Val := <-bufferedChan2:
			fmt.Println(c2Val)
		default:
			done = true
		}
	}

}
