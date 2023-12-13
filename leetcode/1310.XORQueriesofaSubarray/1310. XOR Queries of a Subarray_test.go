package leetcode

import (
	"reflect"
	"testing"
)

func Test_xorQueries(t *testing.T) {
	tests := []struct {
		name    string
		arr     []int
		queries [][]int
		want    []int
	}{
		{
			name:    "Example 1",
			arr:     []int{1, 3, 4, 8},
			queries: [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}},
			want:    []int{2, 7, 14, 8},
		},
		{
			name:    "Example 2",
			arr:     []int{4, 8, 2, 10},
			queries: [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}},
			want:    []int{8, 0, 4, 4},
		},
		// Add more test cases here...
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xorQueries(tt.arr, tt.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("xorQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
