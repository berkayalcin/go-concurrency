package main

import "fmt"

func main() {
	// Blocking read and write
	bufferedChan := make(chan int, 5)

	bufferedChan <- 1
	bufferedChan <- 2
	bufferedChan <- 3
	bufferedChan <- 4
	bufferedChan <- 5
	bufferedChan <- 6
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)
}
