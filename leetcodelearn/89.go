package leetcodelearn

func GrayCode(n int) []int {

	var r []int
	total := 1 << n
	used := make([]bool, total)

	var backtrack func() bool

	backtrack = func() bool {
		if len(r) == total {
			return true
		}
		for i := 0; i < total; i++ {
			if used[i] {
				continue
			}
			if !isValid1(r, i, n) {
				continue
			}

			used[i] = true

			r = append(r, i)
			if backtrack() {
				return true
			}
			r = r[:len(r)-1]
			used[i] = false
		}
		return false
	}
	backtrack()
	return r
}

func isValid1(r []int, i int, n int) bool {
	if len(r) == 0 {
		return true
	}
	last := r[len(r)-1]
	if !hasOnlyOneDifferentBit2(last, i) {
		return false
	}
	if len(r) == (1<<n)-1 {
		if !hasOnlyOneDifferentBit2(r[0], i) {
			return false
		}
	}
	return true
}

func hasOnlyOneDifferentBit2(a, b int) bool {
	xorResult := a ^ b
	return xorResult != 0 && (xorResult&(xorResult-1)) == 0
}
