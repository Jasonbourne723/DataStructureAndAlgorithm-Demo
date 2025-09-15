package leetcodelearn

func MaxProfit(prices []int, strategy []int, k int) int64 {
	var t int
	var max int64
	var total int64

	for i := 0; i < len(prices); i++ {
		total += int64(prices[i] * strategy[i])
	}

	l, m, r := 0, k/2, k-1
	for i := 0; i < k; i++ {
		if i < k/2 {
			t -= strategy[i] * prices[i]
		} else {
			t += prices[i] - (strategy[i] * prices[i])
		}
	}
	if int64(t) > 0 {
		max = int64(t)
	}

	for r < len(prices)-1 {

		t += strategy[l] * prices[l]
		l++

		t -= prices[m]
		m++

		r++
		t += prices[r] - (strategy[r] * prices[r])

		if int64(t) > max {
			max = int64(t)
		}
	}
	return max + total
}
