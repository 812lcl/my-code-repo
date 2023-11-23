package leetcode

func sortColors(nums []int) {
	red, white := 0, 0
	for _, num := range nums {
		if num == 0 {
			red++
		} else if num == 1 {
			white++
		}
	}
	for i := 0; i < len(nums); i++ {
		if red > 0 {
			nums[i] = 0
			red--
		} else if white > 0 {
			nums[i] = 1
			white--
		} else {
			nums[i] = 2
		}
	}
}
