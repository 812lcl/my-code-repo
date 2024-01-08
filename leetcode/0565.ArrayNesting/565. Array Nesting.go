package leetcode

func arrayNesting(nums []int) int {
	uniq := make(map[int]bool)
	var max int
	for i := 0; i < len(nums); i++ {
		var length int
		for j := nums[i]; !uniq[j]; {
			length++
			uniq[j] = true
			j = nums[j]
		}
		clear(uniq)
		if length > max {
			max = length
		}
	}
	return max
}

func arrayNesting1(nums []int) int {
	maxSize := 0
	for i := 0; i < len(nums); i++ {
		size := 0
		for j := i; nums[j] >= 0; size++ {
			nums[j], j = -1, nums[j]
		}
		if size > maxSize {
			maxSize = size
		}
	}
	return maxSize
}
