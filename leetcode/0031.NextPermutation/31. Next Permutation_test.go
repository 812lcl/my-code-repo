package leetcode

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3},
			want: []int{1, 3, 2},
		},
		{
			name: "Example 2",
			nums: []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
		{
			name: "Example 3",
			nums: []int{1, 1, 5},
			want: []int{1, 5, 1},
		},
		{
			name: "Example 4",
			nums: []int{1},
			want: []int{1},
		},
		{
			name: "Example 5",
			nums: []int{8, 9, 6, 10, 7, 2},
			want: []int{8, 9, 7, 2, 6, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("nextPermutation() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
