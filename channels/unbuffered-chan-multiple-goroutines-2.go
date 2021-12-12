package main

import (
	"fmt"
	"time"
)

func main() {
	// Blocking read and write
	unbufferedChan := make(chan int)

	go func(unbufferedChan <-chan int) {
		// Blocks until data arrives
		value := <-unbufferedChan
		fmt.Println(value)
	}(unbufferedChan)

	go func(unbufferedChan chan int) {
		// Blocks until data arrives
		unbufferedChan <- 1
	}(unbufferedChan)

	// now scheduler have time to schedule
	time.Sleep(time.Second)

	fmt.Println("Hello Channels")
}
