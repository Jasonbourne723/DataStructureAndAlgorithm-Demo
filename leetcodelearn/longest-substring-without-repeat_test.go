package leetcodelearn

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	var length = LengthOfLongestSubstring("abcabc")

	if length != 3 {
		t.Errorf("无重复字符错误")
	}
}
