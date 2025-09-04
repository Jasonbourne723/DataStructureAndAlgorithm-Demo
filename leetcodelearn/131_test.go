package leetcodelearn

import "testing"

func TestPartition(t *testing.T) {
	s := "aab"

	r := Partition(s)

	for i := 0; i < len(r); i++ {
		print(r[i])
	}

}
