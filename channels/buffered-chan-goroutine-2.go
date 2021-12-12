package main

import (
	"fmt"
)

func main() {
	// Blocking read and write
	bufferedChan := make(chan int, 5)
	done := make(chan bool)
	go func(bufferedChannel chan int) {
		for val := range bufferedChan {
			fmt.Println(val)
		}
		fmt.Println("Channel closed")
		done <- true
	}(bufferedChan)

	bufferedChan <- 1
	bufferedChan <- 2
	bufferedChan <- 3
	bufferedChan <- 4
	bufferedChan <- 5
	bufferedChan <- 6
	bufferedChan <- 7
	bufferedChan <- 8
	bufferedChan <- 9

	close(bufferedChan)

	<-done
	fmt.Println("Main Finished")
}
