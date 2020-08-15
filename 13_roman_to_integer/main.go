package leetcode

func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	var maps = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	var total, last int
	for i, _ := range s {
		char := s[len(s)-(i+1) : len(s)-i]
		var n = maps[char]
		if n < last {
			total = total - n
		} else {
			total = total + n
		}

		last = n
	}
	return total
}
