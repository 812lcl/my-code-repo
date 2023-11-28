package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				if j-i <= k {
					return true
				}
				continue
			}
		}
	}
	return false
}

// sliding window
func containsNearbyDuplicate1(nums []int, k int) bool {
	if len(nums) <= 1 {
		return false
	}
	if k <= 0 {
		return false
	}
	record := make(map[int]bool, len(nums))
	for i, n := range nums {
		if _, found := record[n]; found {
			return true
		}
		record[n] = true
		if len(record) == k+1 {
			delete(record, nums[i-k])
		}
	}
	return false
}
