package leetcode

import (
	"testing"
)

func Test_findUnsortedSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "case 12",
			nums: []int{-1, -1, -1, -1, -1},
			want: 0,
		},
		{
			name: "case 11",
			nums: []int{2, 3, 1, 4, 5},
			want: 3,
		},
		{
			name: "case 10",
			nums: []int{1, 3, 5, 4, 2},
			want: 4,
		},
		{
			name: "case 0",
			nums: []int{1, 3, 2, 2, 2},
			want: 4,
		},
		{
			name: "case 7",
			nums: []int{2, 3, 3, 2, 4},
			want: 3,
		},
		{
			name: "case 8",
			nums: []int{1, 2, 3, 3, 3},
			want: 0,
		},
		{
			name: "case 9",
			nums: []int{1, 2, 4, 5, 3},
			want: 3,
		},
		{
			name: "case 1",
			nums: []int{1, 2, 3, 4, 5},
			want: 0,
		},
		{
			name: "case 2",
			nums: []int{2, 1, 3, 4, 5},
			want: 2,
		},
		{
			name: "case 3",
			nums: []int{1, 2, 5, 4, 3},
			want: 3,
		},
		{
			name: "case 4",
			nums: []int{1, 3, 2, 4, 5},
			want: 2,
		},
		{
			name: "case 5",
			nums: []int{1, 2, 3, 5, 4},
			want: 2,
		},
		{
			name: "case 6",
			nums: []int{5, 4, 3, 2, 1},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findUnsortedSubarray(tt.nums); got != tt.want {
				t.Errorf("findUnsortedSubarray() = %v, want %v", got, tt.want)
			}
			if got := findUnsortedSubarray1(tt.nums); got != tt.want {
				t.Errorf("findUnsortedSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
