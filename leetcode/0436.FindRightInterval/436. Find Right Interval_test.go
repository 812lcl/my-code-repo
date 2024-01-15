package leetcode

import (
	"reflect"
	"testing"
)

func Test_findRightInterval(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      []int
	}{
		{
			name: "Example 1",
			intervals: [][]int{
				{1, 2},
				{2, 3},
				{0, 1},
				{3, 4},
			},
			want: []int{1, 3, 0, -1},
		},
		{
			name: "Example 2",
			intervals: [][]int{
				{3, 4},
				{2, 3},
				{1, 2},
			},
			want: []int{-1, 0, 1},
		},
		{
			name: "Example 3",
			intervals: [][]int{
				{1, 1},
				{3, 4},
			},
			want: []int{0, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRightInterval(tt.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRightInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
