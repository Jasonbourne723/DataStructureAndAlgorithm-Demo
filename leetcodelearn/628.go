package leetcodelearn

import (
	"math"
	"sort"
)

func maximumProduct(nums []int) int {
	min1, min2 := math.MaxInt, math.MaxInt
	max1, max2, max3 := math.MinInt, math.MinInt, math.MinInt

	for _, v := range nums {
		if v < min1 {
			min2 = min1
			min1 = v
		} else if v < min2 {
			min2 = v
		}

		if v > max1 {
			max3 = max2
			max2 = max1
			max1 = v
		} else if v > max2 {
			max3 = max2
			max2 = v
		} else if v > max3 {
			max3 = v
		}
	}

	return max(max1*max2*max3, min1*min2*max1)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func findErrorNums(nums []int) []int {
	ans := make([]int, 2)
	sort.Ints(nums)

	m := make(map[int]bool, len(nums))

	for i := 0; i < len(nums); i++ {

		if _, exist := m[nums[i]]; exist {
			ans[0] = nums[i]
		}
		m[nums[i]] = true

		if nums[i] != i+1 && ans[1] == 0 {
			ans[1] = i + 1
		}
		if ans[0] != 0 && ans[1] != 0 {
			break
		}
	}
	return ans
}

func removeDuplicates(nums []int) int {

	if len(nums) < 3 {
		return len(nums)
	}

	s := 2
	for f := 2; f < len(nums); f++ {
		if nums[f] != nums[f-2] {
			nums[s] = nums[f]
			s++
		}
	}
	return s
}

func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)

	temp := nums[0] + nums[1] + nums[len(nums)-1]
	if temp == target {
		return temp
	}
	for i := 0; i < len(nums)-2; i++ {
		second := i + 1
		thrid := len(nums) - 1
		for second != thrid {
			sum := nums[i] + nums[second] + nums[thrid]
			if math.Abs(float64(target-temp)) > math.Abs(float64(target-sum)) {
				temp = sum
				if temp == target {
					return temp
				}
			}
			if sum > target {
				thrid--
			} else {
				second++
			}
		}
	}
	return temp
}

func letterCombinations(digits string) []string {
	m := make(map[byte][]byte, 8)
	m['2'] = []byte{'a', 'b', 'c'}
	m['3'] = []byte{'d', 'e', 'f'}
	m['4'] = []byte{'g', 'h', 'i'}
	m['5'] = []byte{'j', 'k', 'l'}
	m['6'] = []byte{'m', 'n', 'o'}
	m['7'] = []byte{'p', 'q', 'r', 's'}
	m['8'] = []byte{'t', 'u', 'v'}
	m['9'] = []byte{'w', 'x', 'y', 'z'}

	dis := [][]byte{}
	if len(digits) == 0 {
		return nil
	}
	for _, v := range m[digits[0]] {
		dis = append(dis, []byte{v})
	}
	for i := 1; i < len(digits); i++ {
		dis = ca(dis, m[digits[i]])
	}
	s := []string{}
	for i, _ := range dis {
		s = append(s, string(dis[i]))
	}

	return s
}

func ca(dis [][]byte, nums []byte) [][]byte {

	n := make([][]byte, 0, len(dis)*len(nums))
	for i, _ := range dis {
		for j, _ := range nums {
			b := append([]byte(nil), dis[i]...)
			b = append(dis[i], nums[j])
			n = append(n, b)
		}
	}

	return n
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	i := 1
	node1 := head
	node2 := head
	var prev *ListNode

	for node1 != nil {
		if i >= n {
			prev = node2
			node2 = node2.Next
		}
		node1 = node1.Next
		i++
	}

	prev.Next = node2.Next
	return head
}

func isValid(s string) bool {

	s1 := stack{
		data: []rune{},
		pos:  -1,
	}
	end := []rune{'}', ']', ')'}
	m := map[rune]rune{'}': '{', ']': '[', ')': '('}

	for _, v := range s {
		if exist(end, v) {
			t, ok := s1.pop()
			if !ok {
				return false
			} else {
				if t != m[v] {
					return false
				}
			}
		} else {
			s1.push(v)
		}
	}
	return true
}

func exist(ans []rune, t rune) bool {
	for _, v := range ans {
		if v == t {
			return true
		}
	}
	return false
}

type stack struct {
	data []rune
	pos  int
}

func (s *stack) pop() (rune, bool) {
	if s.pos < 0 {
		return '0', false
	} else {
		t := s.data[s.pos]
		s.pos--
		return t, true
	}
}

func (s *stack) push(v rune) {
	s.pos++
	s.data = append(s.data, v)
}
