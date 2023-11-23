package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortColors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "case 1",
			nums: []int{2, 0, 2, 1, 1, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "case 2",
			nums: []int{2, 0, 1},
			want: []int{0, 1, 2},
		},
		{
			name: "case 3",
			nums: []int{0},
			want: []int{0},
		},
		{
			name: "case 4",
			nums: []int{1},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("sortColors() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
