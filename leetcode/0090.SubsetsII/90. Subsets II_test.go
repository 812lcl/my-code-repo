package leetcode

import (
	"reflect"
	"testing"
)

func Test_subsetsWithDup(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "case 1",
			nums: []int{1, 2, 2},
			want: [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}},
		},
		{
			name: "case 2",
			nums: []int{0},
			want: [][]int{{}, {0}},
		},
		{
			name: "case 3",
			nums: []int{1, 2, 3},
			want: [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsetsWithDup(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsetsWithDup() = %v, want %v", got, tt.want)
			}
		})
	}
}
