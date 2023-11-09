package leetcode

func removeDuplicates(nums []int) int {
	len := len(nums)
	var i, j int
	for i = 0; i < len-1; {
		for ; j < len; j++ {
			if nums[i] == nums[j] {
				continue
			}
			nums[i+1] = nums[j]
			i++
			break
		}
		if j == len {
			break
		}
	}
	nums = nums[:i+1]
	return i + 1
}
