package leetcode

import "math"

func findUnsortedSubarray(nums []int) int {
	left, right, max, min := -1, -1, math.MinInt, math.MaxInt
	for i := 0; i < len(nums)-1; i++ {
		if left != -1 && nums[i] > max {
			max = nums[i]
		}
		if left != -1 && nums[i+1] < min {
			min = nums[i+1]
		}
		if nums[i+1] < max {
			right = i + 1
			continue
		}
		if nums[i] <= nums[i+1] {
			continue
		}
		if left == -1 {
			left = i
		}
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i+1] < min {
			min = nums[i+1]
		}
		right = i + 1
	}
	if left == right {
		return 0
	}
	for i := left - 1; i >= 0 && nums[i] > min; i-- {
		left = i
	}
	return right - left + 1
}

func findUnsortedSubarray1(nums []int) int {
	n, left, right, minR, maxL, isSort := len(nums), -1, -1, math.MaxInt32, math.MinInt32, false
	// left
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			isSort = true
		}
		if isSort {
			minR = min(minR, nums[i])
		}
	}
	isSort = false
	// right
	for i := n - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			isSort = true
		}
		if isSort {
			maxL = max(maxL, nums[i])
		}
	}
	// minR
	for i := 0; i < n; i++ {
		if nums[i] > minR {
			left = i
			break
		}
	}
	// maxL
	for i := n - 1; i >= 0; i-- {
		if nums[i] < maxL {
			right = i
			break
		}
	}
	if left == -1 || right == -1 {
		return 0
	}
	return right - left + 1
}
