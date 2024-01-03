package sort

import "fmt"

func BubbleSort(nums []int) []int {
	n := len(nums)
	flag := false
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
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
