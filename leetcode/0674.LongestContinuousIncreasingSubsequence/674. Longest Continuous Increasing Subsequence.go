package leetcode

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := 1
	for i := 0; i < len(nums)-1; i++ {
		length := 1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[j-1] {
				length++
				if length > max {
					max = length
				}
			} else {
				i = j - 1
				break
			}
		}
	}
	return max
}
