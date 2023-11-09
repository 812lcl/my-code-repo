package sort

import "fmt"

func InsertSort(nums []int) []int {
	a := 0
	b := len(nums) - 1
	for i := a + 1; i <= b; i++ {
		for j := i; j > a && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]

		}
	}
	fmt.Println(nums)
	return nums
}
