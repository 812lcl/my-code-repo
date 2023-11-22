package leetcode

import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	for i, j := 0, len(str)-1; i <= j; {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isPalindrome1(x int) bool {
	var b int
	var t int = x
	if x < 0 {
		return false
	}
	for x > 0 {
		b = b*10 + x%10
		x = x / 10

	}
	if t == b {
		return true
	} else {
		return false
	}

}
