package errortest

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrorIs(t *testing.T) {

	err := errors.New("failed")

	err1 := fmt.Errorf("err1 is : %w ", err)

	err2 := fmt.Errorf("err2 is %w", err1)

	fmt.Printf("err2 is err1: %v\n", errors.Is(err2, err1))
	fmt.Printf("err1 is err2: %v\n", errors.Is(err1, err2))

	fmt.Printf("err2 is err: %v\n", errors.Is(err2, err))
	fmt.Printf("err1 is err: %v\n", errors.Is(err1, err))

}

func TestErrorAs(t *testing.T) {
	err := MyError{Message: "failed"}

	err1 := fmt.Errorf("err1 is : %w ", err)

	err2 := fmt.Errorf("err2 is %w", err1)

	fmt.Printf("err1: %v\n", err1)
	fmt.Printf("err2: %v\n", err2)

	var err3 MyError

	fmt.Printf("err2 as err: %v\n", errors.As(err2, &err3))
	fmt.Printf("err1 as err: %v\n", errors.As(err1, &err3))

	fmt.Printf("err: %v\n", err3)

	fmt.Printf("errors.Unwrap(err2): %v\n", errors.Unwrap(err2))
	fmt.Printf("errors.Unwrap(err1): %v\n", errors.Unwrap(err1))

	switch err2.(type) {
	case *MyError:
		return
	default:
		return
	}

}

type MyError struct {
	Message string
}

func (e MyError) Error() string {
	return e.Message
}
