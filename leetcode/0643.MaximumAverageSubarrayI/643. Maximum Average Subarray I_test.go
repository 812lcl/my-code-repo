package leetcode

import (
	"testing"
)

func Test_findMaxAverage(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		k      int
		result float64
	}{
		{
			name:   "case 1",
			nums:   []int{1, 12, -5, -6, 50, 3},
			k:      4,
			result: 12.75,
		},
		{
			name:   "case 2",
			nums:   []int{3, -1, 2, 4, -2, 5},
			k:      2,
			result: 3,
		},
		{
			name:   "case 3",
			nums:   []int{0, 0, 0, 0, 0},
			k:      3,
			result: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findMaxAverage(tt.nums, tt.k)
			if res != tt.result {
				t.Errorf("findMaxAverage() = %v, want %v", res, tt.result)
			}
		})
	}
}
