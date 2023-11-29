package leetcode

func maxArea(height []int) int {
	max := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			area := (j - i) * min(height[i], height[j])
			if area > max {
				max = area
			}
		}
	}
	return max
}

func maxArea1(height []int) int {
	max := 0
	for i, j := 0, len(height)-1; i < j; {
		area := (j - i) * min(height[i], height[j])
		if area > max {
			max = area
		}
		if height[i] >= height[j] {
			j--
		} else {
			i++
		}
	}
	return max
}
