package leetcode

import "testing"

func Test_findNumberOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "case 1",
			nums: []int{1, 3, 5, 4, 7},
			want: 2,
		},
		{
			name: "case 2",
			nums: []int{2, 2, 2, 2, 2},
			want: 5,
		},
		{
			name: "case 3",
			nums: []int{5, 4, 3, 2, 1},
			want: 5,
		},
		{
			name: "case 4",
			nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumberOfLIS(tt.nums); got != tt.want {
				t.Errorf("findNumberOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
