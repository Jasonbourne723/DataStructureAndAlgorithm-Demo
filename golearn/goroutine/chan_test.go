package main

import (
	"fmt"
	"runtime"
	"testing"
)

func TestRuntime(t *testing.T) {

	go func(s string) {
		for i := 0; i < 10; i++ {
			fmt.Printf("s: %v\n", s)
		}
	}("world")

	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println("hello")
	}
}

func TestChan(t *testing.T) {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {

		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for {
			if d, ok := <-ch1; ok {
				fmt.Printf("d: %v\n", d)
			} else {
				break
			}
		}
		close(ch2)
	}()
	<-ch2
}

