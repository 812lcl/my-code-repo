package leetcode

import "812lcl/my-code-repo/sort"

func sortArray(nums []int) []int {
	sort.QuickSort(nums)
	sort.MergeSort(nums)
	sort.SelectSort(nums)
	sort.InsertSort(nums)
	sort.BubbleSort(nums)
	return nums
}
