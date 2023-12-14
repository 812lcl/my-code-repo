package leetcode

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var stack []int
	stack = append(stack, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] > stack[len(stack)-1] {
			stack = append(stack, nums[i])
		} else {
			return i - 1
		}
	}
	return len(nums) - 1
}

func findPeakElement1(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)>>1
		// 如果 mid 较大，则左侧存在峰值，high = m，如果 mid + 1 较大，则右侧存在峰值，low = mid + 1
		if nums[mid] > nums[mid+1] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
