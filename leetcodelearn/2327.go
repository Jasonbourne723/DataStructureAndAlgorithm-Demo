package leetcodelearn

func PeopleAwareOfSecret(n int, delay int, forget int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	c := 1
	q := []int{0}
	for i := 1; i < n; i++ {
		for j := len(q) - 1; j >= 0; j-- {
			q[j]++
			if q[j] >= delay {
				if q[j] == forget {
					q = append(q[:j], q[j+1:]...)
					c--
				} else {
					q = append(q, 0)
					c++
				}
			}
		}
	}
	c %= (10 ^ 9 + 7)
	return c
}
