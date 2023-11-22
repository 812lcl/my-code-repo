package leetcode

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		want   int
	}{
		{
			name:   "case 1",
			target: 7,
			nums:   []int{2, 3, 1, 2, 4, 3},
			want:   2,
		},
		{
			name:   "case 2",
			target: 11,
			nums:   []int{1, 2, 3, 4, 5},
			want:   3,
		},
		{
			name:   "case 3",
			target: 15,
			nums:   []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8},
			want:   2,
		},
		// Add more test cases here...
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.target, tt.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
