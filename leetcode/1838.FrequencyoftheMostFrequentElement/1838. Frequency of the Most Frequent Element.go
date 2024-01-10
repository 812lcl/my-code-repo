package leetcode

import "sort"

func maxFrequency(nums []int, k int) int {
	max := 1
	sort.Ints(nums)
	for i := len(nums) - 1; i > 0; i-- {
		diff := k
		freq := 1
		for j := i - 1; j >= 0; j-- {
			if diff >= nums[i]-nums[j] {
				diff -= nums[i] - nums[j]
				freq++
				if freq > max {
					max = freq
				}
			}
		}
	}
	return max
}

func maxFrequency1(nums []int, k int) int {
	sort.Ints(nums)

	windowStart := 0
	result, sum := 1, 0
	for windowEnd := 0; windowEnd < len(nums); windowEnd++ {
		sum += nums[windowEnd]

		if nums[windowEnd]*(windowEnd-windowStart+1) > sum+k {
			sum -= nums[windowStart]
			windowStart++
		}

		result = max(result, windowEnd-windowStart+1)
	}
	return result
}
