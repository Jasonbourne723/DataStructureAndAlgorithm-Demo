package mapreducetest

import (
	"fmt"
	"testing"
)

func TestMapReduce(t *testing.T) {

	a := []int{1, 2, 3, 4, 5, 5, 9}

	b := Filter(a, func(t int) bool {
		return t > 3
	})

	fmt.Printf("b: %v\n", b)

}

func MapStringToInt(array []string, option func(s string) int64) []int64 {

	a := make([]int64, len(array))

	for i, v := range array {
		a[i] = option(v)
	}

	return a
}

func ReduceForCount[T any](s []T, fn func(t T) int) int {

	count := 0

	for _, v := range s {
		count += fn(v)
	}
	return count
}

func Filter[T any](s []T, fn func(t T) bool) []T {

	o := make([]T, 0, len(s))

	for _, v := range s {
		if fn(v) {
			o = append(o, v)
		}
	}
	return o
}
