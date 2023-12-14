package leetcode

import "testing"

func Test_findPeakElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Case 1",
			nums:     []int{1, 2, 3, 1},
			expected: 2,
		},
		{
			name:     "Case 3",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "Case 4",
			nums:     []int{5, 4, 3, 2, 1},
			expected: 0,
		},
		{
			name:     "Case 5",
			nums:     []int{1},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findPeakElement(tt.nums)
			if result != tt.expected {
				t.Errorf("findPeakElement(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
			result = findPeakElement1(tt.nums)
			if result != tt.expected {
				t.Errorf("findPeakElement(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}
