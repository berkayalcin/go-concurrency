package main

import "fmt"

func main() {
	// Blocking read and write
	unbufferedChan := make(chan int)

	var unbufferedChan2 chan int

	fmt.Println(unbufferedChan)
	fmt.Println(unbufferedChan2) // nil

	go func(unbufferedChan <-chan int) {
		// Blocks until data arrives
		value := <-unbufferedChan
		fmt.Println(value)
	}(unbufferedChan)

	unbufferedChan <- 1
}
