package leetcode

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	radixSort(nums)
	var res int
	for i, j := 0, 1; i < len(nums)-1; i++ {
		diff := nums[j] - nums[i]
		if diff > res {
			res = diff
		}
		j++
	}
	return res
}

func maximumGap1(nums []int) int {
	if nums == nil || len(nums) < 2 {
		return 0
	}
	// m is the maximal number in nums
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		m = max(m, nums[i])
	}
	exp := 1 // 1, 10, 100, 1000 ...
	R := 10  // 10 digits
	aux := make([]int, len(nums))
	for (m / exp) > 0 { // Go through all digits from LSB to MSB
		count := make([]int, R)
		for i := 0; i < len(nums); i++ {
			count[(nums[i]/exp)%10]++
		}
		for i := 1; i < len(count); i++ {
			count[i] += count[i-1]
		}
		for i := len(nums) - 1; i >= 0; i-- {
			tmp := count[(nums[i]/exp)%10]
			tmp--
			aux[tmp] = nums[i]
			count[(nums[i]/exp)%10] = tmp
		}
		for i := 0; i < len(nums); i++ {
			nums[i] = aux[i]
		}
		exp *= 10
	}
	maxValue := 0
	for i := 1; i < len(aux); i++ {
		maxValue = max(maxValue, aux[i]-aux[i-1])
	}
	return maxValue
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func radixSort(nums []int) {
	if len(nums) < 2 {
		return
	}

	// Find the maximum number in the array
	maxNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}

	// Perform counting sort for every digit
	exp := 1
	for maxNum/exp > 0 {
		countSort(nums, exp)
		exp *= 10
	}
}

func countSort(nums []int, exp int) {
	n := len(nums)
	output := make([]int, n)
	count := make([]int, 10)

	// Count the occurrences of each digit
	for i := 0; i < n; i++ {
		count[(nums[i]/exp)%10]++
	}

	// Calculate the cumulative count
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// Build the output array
	for i := n - 1; i >= 0; i-- {
		output[count[(nums[i]/exp)%10]-1] = nums[i]
		count[(nums[i]/exp)%10]--
	}

	// Copy the output array to the original array
	for i := 0; i < n; i++ {
		nums[i] = output[i]
	}
}
