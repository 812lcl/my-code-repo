package leetcode

import "testing"

func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "example 1",
			nums:   []int{-1, 2, 1, -4},
			target: 1,
			want:   2,
		},
		{
			name:   "example 2",
			nums:   []int{0, 0, 0},
			target: 1,
			want:   0,
		},
		{
			name:   "example 3",
			nums:   []int{1, 1, 1, 0},
			target: -100,
			want:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.nums, tt.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
			if got := threeSumClosest1(tt.nums, tt.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
