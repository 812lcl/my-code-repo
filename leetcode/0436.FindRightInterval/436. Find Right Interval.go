package leetcode

import "math"

func findRightInterval(intervals [][]int) []int {
	var res []int
	for i := 0; i < len(intervals); i++ {
		min := math.MaxInt
		val := -1
		for j := 0; j < len(intervals); j++ {
			if intervals[j][0] >= intervals[i][1] && intervals[j][0] < min {
				min = intervals[j][0]
				val = j
			}
		}
		res = append(res, val)
	}
	return res
}
