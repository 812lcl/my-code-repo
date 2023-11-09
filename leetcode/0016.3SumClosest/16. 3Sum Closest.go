package leetcode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	min := math.MaxInt
	res := target
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				v := nums[i] + nums[j] + nums[k]
				if v == target {
					return target
				}
				if abs(v-target) < min {
					min = abs(v - target)
					res = v
				}
			}
		}
	}
	return res
}

func threeSumClosest1(nums []int, target int) int {
	length, res, min := len(nums), 0, math.MaxInt
	if length > 2 {
		sort.Ints(nums)
		for i := 0; i < length; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			for j, k := i+1, length-1; j < k; {
				sum := nums[i] + nums[j] + nums[k]
				if abs(sum-target) < min {
					min = abs(sum - target)
					res = sum
				}
				if sum == target {
					return target
				} else if sum > target {
					k--
				} else {
					j++
				}

			}
		}
	}
	return res
}

func abs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}
