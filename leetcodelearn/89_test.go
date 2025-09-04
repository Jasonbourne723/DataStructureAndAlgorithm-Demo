package leetcodelearn

import "testing"

func TestGrayCode(t *testing.T) {

	result := GrayCode(2)
	for i := 0; i < len(result); i++ {
		println(result[i])
	}
}
