package main

import (
	"fmt"
	"sync"
	"testing"
)

var wg1 sync.WaitGroup

func TestWaitGroup(t *testing.T) {

	for i := 0; i < 10; i++ {
		wg1.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("i: %v\n", i)
		}(i)
	}

	wg1.Wait()
	fmt.Println("End ... ")
}
