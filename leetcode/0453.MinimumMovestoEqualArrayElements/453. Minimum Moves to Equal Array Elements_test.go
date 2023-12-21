package leetcode

import (
	"testing"
)

func Test_minMoves(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "case 1",
			nums: []int{1, 2, 3},
			want: 3,
		},
		{
			name: "case 2",
			nums: []int{1, 1, 1},
			want: 0,
		},
		{
			name: "case 3",
			nums: []int{2, 3, 4},
			want: 3,
		},
		{
			name: "case 4",
			nums: []int{5, 5, 5},
			want: 0,
		},
		{
			name: "case 5",
			nums: []int{1, 2, 3, 4, 5},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMoves(tt.nums); got != tt.want {
				t.Errorf("minMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
