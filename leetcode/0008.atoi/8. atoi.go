package leetcode

import "math"

func myAtoi(s string) int {
	var res int
	min, max := math.MinInt32, math.MaxInt32
	var negative, begin bool
	for _, v := range s {
		if (string(v) == " " || string(v) == "-" || string(v) == "+") && !begin {
			if string(v) == "-" {
				negative = true
				begin = true
			} else if string(v) == "+" {
				negative = false
				begin = true
			}
			continue
		} else {
			begin = true
		}

		var n int
		if n = toNum(v); n < 0 || n > 9 {
			break
		}
		res = res*10 + n
		if res > max {
			if negative {
				return min
			} else {
				return max
			}
		}

	}
	if negative {
		return -res
	}
	return res
}

func toNum(v rune) int {
	return int(v - 48)
}
