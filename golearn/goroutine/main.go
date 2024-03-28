package main

import (
	"sync"
)

var wg sync.WaitGroup

func main() {

	ch1 := make(chan int)
	ch1 <- 10
	wg.Add(1)

	go func() {

		ch1 <- 10
		wg.Done()
	}()

	wg.Wait()

}
