package leetcode

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}
	c := 1
	for i := len(digits) - 1; i >= 0; i-- {
		t := digits[i]
		digits[i] = (t + c) % 10
		c = (t + c) / 10
		if c == 0 {
			break
		}
	}

	if c == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
