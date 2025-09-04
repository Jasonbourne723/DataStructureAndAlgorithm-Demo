package leetcodelearn

func ShipWithinDays(weights []int, days int) int {
	var total int
	var max int
	for _, item := range weights {
		total += item
		if item > max {
			max = item
		}
	}
	return bsearch(weights, max, total, days)

}

func bsearch(weights []int, min int, max int, days int) int {
	days_min := getdays(weights, min)
	if days_min <= days {
		return min
	}
	if min > max {
		return max
	}

	middle := (max-min)/2 + min
	days_middle := getdays(weights, middle)

	if days_middle <= days {
		return bsearch(weights, min+1, middle, days)
	} else {
		return bsearch(weights, middle+1, max, days)
	}
}

func getdays(weights []int, load int) int {
	var tmp int
	day := 1
	for _, w := range weights {
		if tmp+w < load {
			tmp += w
		} else if tmp+w == load {
			tmp = 0
			day += 1
		} else {
			tmp = w
			day += 1
		}
	}
	if tmp == 0 {
		day -= 1
	}
	return day
}
