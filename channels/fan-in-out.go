package main

import (
	"fmt"
	"time"
)

func main() {
	// Blocking read and write
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	outputChan := make(chan int, 10000)
	done := make(chan int)

	go func() {
		for {
			fmt.Println("Worker Waits")
			val, open := <-outputChan
			if !open {
				break
			}

			fmt.Println(val)
		}
		done <- 1
	}()

	go func() {
		for {
			c1 <- 1
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			c2 <- 2
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			c3 <- 3
			time.Sleep(time.Second)
		}
	}()

	fanIn([]chan int{c1, c2, c3}, outputChan)
	<-done
}

func fanIn(i []chan int, outputChan chan int) {
	for _, ch := range i {
		go func(c chan int) {
			for {
				val, open := <-c
				if !open {
					break
				}
				outputChan <- val
			}
		}(ch)
	}
}
