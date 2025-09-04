package leetcodelearn

import "testing"

func TestSolveNQueens(t *testing.T) {

	r := SolveNQueens(4)
	for _, item := range r {
		for _, item1 := range item {
			print(item1)
			print(",")
		}
		println("\n")
	}

}
