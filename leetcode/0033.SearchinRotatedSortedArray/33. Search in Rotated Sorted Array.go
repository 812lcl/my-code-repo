package leetcode

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if target < nums[0] && target > nums[n-1] {
		return -1
	}
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[right] {
			if nums[mid] > target && target >= nums[left] {
				right = mid - 1
			} else if target < nums[mid] && target > nums[right] {
				return -1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < nums[left] {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else if target > nums[mid] && target < nums[left] {
				return -1
			} else {
				right = mid - 1
			}
		} else {
			if target > nums[right] || target < nums[left] {
				return -1
			}
			if nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}
