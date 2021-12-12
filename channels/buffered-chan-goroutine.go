package main

import "fmt"

func main() {
	// Blocking read and write
	bufferedChan := make(chan int, 5)

	go func(bufferedChannel chan int) {
		for {
			val := <-bufferedChan
			fmt.Println(val)
		}
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

}
