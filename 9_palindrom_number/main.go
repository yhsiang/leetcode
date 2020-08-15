package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	var o = x
	tmp := 0
	for x != 0 {
		tmp = tmp*10 + x%10
		x = x / 10
	}
	return tmp == o
}
