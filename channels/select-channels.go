package main

import "fmt"

func main() {
	// Blocking read and write
	bufferedChan := make(chan int, 1)
	bufferedChan <- 1

	select {
	case c1Val := <-bufferedChan:
		fmt.Println(c1Val)
	}

}
