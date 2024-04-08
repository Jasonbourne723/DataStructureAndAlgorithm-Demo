package errortest

import (
	//"errors"
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func TestError(t *testing.T) {

	if err := fail(); err != nil {
		fmt.Printf("err: %+v\n", err)
		fmt.Printf("errors.Cause(err): %v\n", errors.Cause(err))
	}

}

func fail() error {

	err := errors.New("fail")

	return errors.Wrap(err, "authorication err: ")
}

type AuthicationError struct {
	err error
}

func (err *AuthicationError) Error() string {
	return fmt.Sprintf("Authication Error is %s", err.err.Error())
}
