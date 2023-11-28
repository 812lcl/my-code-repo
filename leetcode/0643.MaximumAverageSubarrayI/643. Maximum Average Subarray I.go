package leetcode

func findMaxAverage(nums []int, k int) float64 {
	if len(nums) < k {
		return 0
	}
	var sum int
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	avg := float64(sum) / float64(k)
	res := avg
	for i := k; i < len(nums); i++ {
		avg = avg + float64(nums[i]-nums[i-k])/float64(k)
		if avg > res {
			res = avg
		}
	}
	return res
}
