package sort

import "fmt"

func BubbleSort(nums []int) []int {
	n := len(nums)
	flag := false
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	fmt.Println(nums)
	return nums
}
