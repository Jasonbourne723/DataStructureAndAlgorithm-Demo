package leetcodelearn

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{1, 5, 6, 9, 8, 7}
	target := 9

	a := twoSum(nums, target)

	if !reflect.DeepEqual(a, []int{0, 4}) {
		t.Errorf("错误")
	}
}
