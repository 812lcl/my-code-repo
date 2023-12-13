package leetcode

import "testing"

func Test_search(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "empty array",
			nums:   []int{},
			target: 5,
			want:   -1,
		},
		{
			name:   "target not in array",
			nums:   []int{1, 2, 3, 4, 5},
			target: 6,
			want:   -1,
		},
		{
			name:   "target at start of array",
			nums:   []int{4, 5, 6, 7, 1, 2, 3},
			target: 4,
			want:   0,
		},
		{
			name:   "target at end of array",
			nums:   []int{4, 5, 6, 7, 1, 2, 3},
			target: 3,
			want:   6,
		},
		{
			name:   "target in rotated part of array",
			nums:   []int{4, 5, 6, 7, 1, 2, 3},
			target: 1,
			want:   4,
		},
		{
			name:   "target in non-rotated part of array",
			nums:   []int{4, 5, 6, 7, 1, 2, 3},
			target: 6,
			want:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.nums, tt.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
