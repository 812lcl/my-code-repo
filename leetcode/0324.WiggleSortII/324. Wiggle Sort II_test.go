package leetcode

import (
	"reflect"
	"testing"
)

func Test_wiggleSort(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "case 1",
			nums: []int{1, 5, 1, 1, 6, 4},
			want: []int{1, 4, 1, 5, 1, 6},
		},
		{
			name: "case 2",
			nums: []int{1, 3, 2, 2, 3, 1},
			want: []int{1, 2, 1, 3, 2, 3},
		},
		{
			name: "case 3",
			nums: []int{1, 1, 1, 1, 1, 1},
			want: []int{1, 1, 1, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wiggleSort(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("wiggleSort() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
