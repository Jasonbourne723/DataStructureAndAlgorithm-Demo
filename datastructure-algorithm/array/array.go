package array

func reverse(a []int) []int {

	if len(a) == 0 || len(a) == 1 {
		return a
	}
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
	return a
}
