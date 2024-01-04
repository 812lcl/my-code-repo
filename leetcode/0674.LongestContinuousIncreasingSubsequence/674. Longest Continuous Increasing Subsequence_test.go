package leetcode

import "testing"

func TestFindLengthOfLCIS(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 3, 5, 4, 7}, 3},
		{[]int{2, 2, 2, 2, 2}, 1},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{5, 4, 3, 2, 1}, 1},
		{[]int{2, 1, 3}, 2},
		{[]int{1}, 1},
		{[]int{}, 0},
	}

	for _, test := range tests {
		result := findLengthOfLCIS(test.nums)
		if result != test.expected {
			t.Errorf("Expected %d, but got %d for input %v", test.expected, result, test.nums)
		}
	}
}
