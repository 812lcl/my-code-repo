package leetcode

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	return binarySearch(nums, target, 0, len(nums)-1)
}

func binarySearch(nums []int, target, left, right int) int {
	if right == left {
		if nums[left] >= target {
			return left
		} else {
			return left + 1
		}
	}

	if nums[left] >= target {
		return left
	}
	if nums[right] < target {
		return right + 1
	}

	mid := (right + left) / 2
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return binarySearch(nums, target, left, mid-1)
	} else {
		return binarySearch(nums, target, mid+1, right)
	}
}
