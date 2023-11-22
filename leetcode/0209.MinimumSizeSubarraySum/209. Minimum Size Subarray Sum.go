package leetcode

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right, res, sum := 0, 0, len(nums)+1, 0
	for right < len(nums) {
		if sum+nums[right] >= target {
			if res > right-left+1 {
				res = right - left + 1
			}
			sum -= nums[left]
			left++
		} else {
			sum += nums[right]
			right++
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}
