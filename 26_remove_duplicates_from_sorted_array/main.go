package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var temp = nums
	var a = []int{}
	var maps = map[int]struct{}{}
	for _, n := range temp {
		if _, ok := maps[n]; !ok {
			maps[n] = struct{}{}
			a = append(a, n)
		}

	}
	copy(nums, a)
	return len(a)
}
