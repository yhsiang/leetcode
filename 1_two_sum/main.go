package leetcode

func twoSum(nums []int, target int) []int {
	var maps = make(map[int]int)
	for i, n := range nums {
		var c = target - n
		if v, ok := maps[c]; ok {
			return []int{v, i}
		}
		maps[n] = i
	}
	return []int{}
}
