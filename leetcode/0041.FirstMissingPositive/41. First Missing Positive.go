package leetcode

func firstMissingPositive(nums []int) int {
	n := len(nums)

	// 第一次遍历，将小于等于 0 和大于 n 的值都设为 n+1
	for i := 0; i < n; i++ {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = n + 1
		}
	}

	// 第二次遍历，将每个元素对应的位置标记为负数
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			nums[num-1] = -abs(nums[num-1])
		}
	}

	// 第三次遍历，找到第一个正数的位置，返回对应的正整数
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

	// 如果数组中都是正整数，返回 n+1
	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
