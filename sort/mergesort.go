package sort

import "fmt"

func MergeSort(nums []int) []int {
	left := 0
	right := len(nums) - 1
	mergeSoft(nums, left, right)
	fmt.Println(nums)
	return nums
}

func mergeSoft(nums []int, left, right int) {
	if left >= right {
		return
	}
	middle := (right-left)/2 + left
	mergeSoft(nums, left, middle)
	mergeSoft(nums, middle+1, right)
	merge(nums, left, middle, right)
}

func merge(nums []int, left, middle, right int) {
	tmp := make([]int, right-left+1)
	i, j, k := left, middle+1, 0
	for i <= middle && j <= right {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
		k++
	}
	for i <= middle {
		tmp[k] = nums[i]
		i++
		k++
	}
	for j <= right {
		tmp[k] = nums[j]
		j++
		k++
	}
	for k := 0; k < len(tmp); k++ {
		nums[left+k] = tmp[k]
	}
}
