package leetcode

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	// dp[i] 代表抢 nums[0...i] 房子的最大价值
	dp := make([]int, n)
	dp[0], dp[1] = nums[0], max(nums[1], nums[0])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	return dp[n-1]
}

func rob1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	preMax, curMax := nums[0], max(nums[1], nums[0])
	for i := 2; i < n; i++ {
		tmp := curMax
		curMax = max(curMax, nums[i]+preMax)
		preMax = tmp
	}
	return curMax
}
