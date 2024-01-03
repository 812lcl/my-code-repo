package sort

import "fmt"

func SelectSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		k := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[k] {
				k = j
			}
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	fmt.Println(nums)
	return nums
}
