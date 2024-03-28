package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {

	ctx := context.Background()

	if deadline, ok := ctx.Deadline(); ok {
		fmt.Printf("deadline: %v\n", deadline)
	} else {
		fmt.Println("don`t has deadline")
	}

	if err := ctx.Err(); err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Println("no error")
	}

	cancelCtx, cancelFunc := context.WithCancel(ctx)

	valueCtx := context.WithValue(cancelCtx, "a", "vb")

	cancelFunc()

	select {
	case <-cancelCtx.Done():
		fmt.Println("cancelCtx is done")
	default:
		fmt.Println("timeout")
	}

	select {
	case <-cancelCtx.Done():
		fmt.Println("cancelCtx is done")
	default:
		fmt.Println("timeout")
	}

	fmt.Printf("valueCtx.Value(\"a\"): %v\n", valueCtx.Value("a"))

}

func TestTimeoutContext(t *testing.T) {
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)

	cancelFunc()

	go func() {
		<-timeoutCtx.Done()
		fmt.Println("timeout" + timeoutCtx.Err().Error())
	}()
}

func TestSelect(t *testing.T) {

	select {
	case <-time.After(2 * time.Second):
		fmt.Print("timeout")
	default:
		fmt.Print("default")
	}
	
}
