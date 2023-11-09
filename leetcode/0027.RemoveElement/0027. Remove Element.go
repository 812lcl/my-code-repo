package leetcode

func removeElement(nums []int, val int) int {
	counter := make(map[int]int)
	for _, v := range nums {
		counter[v]++
	}

	if counter[val] == 0 {
		return len(nums)
	}

	j := 0
	for k, v := range counter {
		if k == val {
			continue
		}
		for i := 0; i < v; i++ {
			nums[j] = k
			j++
		}
	}
	return len(nums) - counter[val]
}

func removeElement1(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
	return j
}
