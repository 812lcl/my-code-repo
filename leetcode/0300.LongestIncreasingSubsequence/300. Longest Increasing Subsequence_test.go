package leetcode

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "case 1",
			nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			want: 4,
		},
		{
			name: "case 2",
			nums: []int{0, 1, 0, 3, 2, 3},
			want: 4,
		},
		{
			name: "case 3",
			nums: []int{7, 7, 7, 7, 7, 7, 7},
			want: 1,
		},
		{
			name: "case 4",
			nums: []int{1, 3, 2, 4, 5, 6, 7, 8, 9},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
			if got := lengthOfLIS1(tt.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
