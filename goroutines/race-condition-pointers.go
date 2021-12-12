package main

import (
	"fmt"
	"sync"
)

type RaceTest struct {
	Val int
}

func main() {
	raceTest := &RaceTest{}
	wg := &sync.WaitGroup{}
	wg.Add(10000)
	mutex := sync.Mutex{}
	for i := 0; i < 10000; i++ {
		go increment(raceTest, wg, &mutex)
	}

	wg.Wait()
	fmt.Println(raceTest)
}

func increment(rt *RaceTest, wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	rt.Val += 1
	mutex.Unlock()
	wg.Done()
}
