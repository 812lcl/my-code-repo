package leetcode

import (
	"testing"
)

func Test_maximumGap(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "case 1",
			nums: []int{5, 2, 3, 1},
			want: 2,
		},
		{
			name: "case 2",
			nums: []int{5, 1, 1, 2, 0, 0},
			want: 3,
		},
		{
			name: "case 3",
			nums: []int{1},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumGap(tt.nums); got != tt.want {
				t.Errorf("maximumGap() = %v, want %v", got, tt.want)
			}
			if got := maximumGap1(tt.nums); got != tt.want {
				t.Errorf("maximumGap() = %v, want %v", got, tt.want)
			}
		})
	}
}
