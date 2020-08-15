package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var bitSet [256]uint8
	result, left, right := 0, 0, 0
	for left < len(s) {
		if right < len(s) && bitSet[s[right]] == 0 {
			bitSet[s[right]] = 1
			right++
		} else {
			bitSet[s[left]] = 0
			left++
		}
		if result < right-left {
			result = right - left
		}
	}
	return result
}
