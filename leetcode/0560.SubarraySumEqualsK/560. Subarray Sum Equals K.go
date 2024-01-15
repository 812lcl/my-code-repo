package leetcode

func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	var count int
	for i := 0; i < len(nums); i++ {
		if nums[i] == k {
			count++
		}
		var sum int
		sum += nums[i]
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}

func subarraySum1(nums []int, k int) int {
	// the hashmap prefix stores the prefix sum from 0 to i and the number of its occurence.
	prefix := map[int]int{0: 1}
	var sum int
	var count int

	for _, num := range nums {
		// sum represents the sum of nums[:i](including nums[i])
		sum += num
		// since the hashmap prefix has stored the sum of the subarray nums[:j](0<=j<=i) and
		// the number of its occurence, so the number of the subarray nums[j:i](including i
		// and j) whose sum is k is equal to the number of the prefix nums[:j] whose sum is
		// sum-k.
		count += prefix[sum-k]
		prefix[sum]++
	}

	return count
}
