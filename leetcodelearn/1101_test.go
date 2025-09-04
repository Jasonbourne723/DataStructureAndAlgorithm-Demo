package leetcodelearn

import "testing"

func TestShipWithinDays(t *testing.T) {
	weights := []int{3, 3, 3, 3, 3, 3}
	days := 2
	min := ShipWithinDays(weights, days)
	println(min)
}
