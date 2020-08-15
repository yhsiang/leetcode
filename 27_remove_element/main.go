package leetcode

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	var temp = nums
	var a = []int{}
	for _, n := range temp {
		if n != val {
			a = append(a, n)
		}

	}
	copy(nums, a)
	return len(a)
}
