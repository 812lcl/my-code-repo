package leetcode

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{2, 1, 1, 2}, 4},
	}

	for _, test := range tests {
		result := rob(test.nums)
		if result != test.expected {
			t.Errorf("Expected %d, but got %d for input %v", test.expected, result, test.nums)
		}
		result = rob1(test.nums)
		if result != test.expected {
			t.Errorf("Expected %d, but got %d for input %v", test.expected, result, test.nums)
		}
	}
}
