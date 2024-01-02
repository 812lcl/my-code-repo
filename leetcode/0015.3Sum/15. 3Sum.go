package leetcode

import (
	"slices"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					s := []int{nums[i], nums[j], nums[k]}
					flag := false
					for _, s1 := range res {
						if slices.Equal(s1, s) {
							flag = true
						}
					}
					if flag {
						continue
					}
					res = append(res, s)
					flag = false
				}
			}
		}
	}
	return res
}

func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	result, start, end, index, sum, length := make([][]int, 0), 0, 0, 0, 0, len(nums)
	for index = 1; index < length-1; index++ {
		start, end = 0, length-1
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		for start < index && end > index {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			sum = nums[start] + nums[end] + nums[index]
			if sum == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if sum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}

func threeSum3(nums []int) [][]int {
	res := [][]int{}
	counter := map[int]int{}
	for _, value := range nums {
		counter[value]++
	}

	uniqNums := []int{}
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums)

	for i := 0; i < len(uniqNums); i++ {
		if (uniqNums[i]*3 == 0) && counter[uniqNums[i]] >= 3 {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i]})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			if (uniqNums[i]*2+uniqNums[j] == 0) && counter[uniqNums[i]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if (uniqNums[j]*2+uniqNums[i] == 0) && counter[uniqNums[j]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			c := 0 - uniqNums[i] - uniqNums[j]
			if c > uniqNums[j] && counter[c] > 0 {
				res = append(res, []int{uniqNums[i], uniqNums[j], c})
			}
		}
	}
	return res
}

func threeSum4(nums []int) [][]int {
	sort.Ints(nums)
	result, length := make([][]int, 0), len(nums)
	for start := 0; start < length-2; start++ {
		if start > 0 && nums[start] == nums[start-1] {
			continue
		}
		for mid, end := start+1, length-1; mid < end; {
			if mid > start+1 && nums[mid] == nums[mid-1] {
				mid++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			sum := nums[start] + nums[mid] + nums[end]
			if sum == 0 {
				result = append(result, []int{nums[start], nums[mid], nums[end]})
				mid++
				end--
			} else if sum > 0 {
				end--
			} else {
				mid++
			}

		}
	}
	return result
}
