package leetcode

func mySqrt(x int) int {
	var res int
	for i := 0; i*i <= x; i++ {
		res = i
	}
	return res
}

func mySqrt1(x int) int {
	l, r := 0, x
	for l < r {
		mid := (l + r + 1) / 2
		if mid*mid <= x {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}

func mySqrt2(x int) int {
	res := x
	for res*res > x {
		res = (res + x/res) / 2
	}
	return res
}
