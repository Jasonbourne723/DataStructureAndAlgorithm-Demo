package leetcodelearn

import (
	"testing"
)

func TestExist(t *testing.T) {

	s3 := [][]byte{{'a', 'A', 'a', 'a', 'a', 'a'}, {'a', 'a', 'A', 'A', 'a', 'A'}, {'A', 'A', 'a', 'A', 'A', 'A'}, {'a', 'A', 'a', 'a', 'a', 'A'}, {'a', 'a', 'a', 'A', 'A', 'a'}, {'a', 'a', 'A', 'A', 'a', 'a'}}
	exist := Exist(s3, "AAaaAaaaaaAaAa")
	print(exist)
}
