package leetcode

func findPrefixScore(nums []int) []int64 {
	if len(nums) == 0 {
		return []int64{}
	}
	conver, score := make([]int64, len(nums)), make([]int64, len(nums))
	var max int
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		conver[i] = int64(nums[i] + max)
		if i == 0 {
			score[i] = conver[i]
		} else {
			score[i] = score[i-1] + conver[i]
		}
	}
	return score
}
