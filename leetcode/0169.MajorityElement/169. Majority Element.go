package leetcode

func majorityElement(nums []int) int {
	counter := make(map[int]int)
	for _, v := range nums {
		counter[v]++
		if counter[v] > len(nums)/2 {
			return v
		}

	}
	return 0
}

func majorityElement1(nums []int) int {
	res, count := nums[0], 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			res, count = nums[i], 1
		} else {
			if nums[i] == res {
				count++
			} else {
				count--
			}
		}
	}
	return res
}
