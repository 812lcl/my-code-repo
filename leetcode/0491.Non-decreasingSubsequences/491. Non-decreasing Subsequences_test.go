package leetcode

import (
	"reflect"
	"testing"
)

func Test_findSubsequences(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Example 1",
			nums: []int{4, 6, 7, 7},
			want: [][]int{{4, 6}, {4, 6, 7}, {4, 6, 7, 7}, {4, 7}, {4, 7, 7}, {6, 7}, {6, 7, 7}, {7, 7}},
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 3, 4},
			want: [][]int{{1, 2}, {1, 2, 3}, {1, 2, 3, 4}, {1, 2, 4}, {1, 3}, {1, 3, 4}, {1, 4}, {2, 3}, {2, 3, 4}, {2, 4}, {3, 4}},
		},
		// Add more test cases here...
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubsequences(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubsequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
