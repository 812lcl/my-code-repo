package leetcode

import (
	"reflect"
	"testing"
)

func Test_findMinHeightTrees(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  []int
	}{
		{
			name:  "example 2",
			n:     6,
			edges: [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}},
			want:  []int{3, 4},
		},
		{
			name:  "example 1",
			n:     4,
			edges: [][]int{{1, 0}, {1, 2}, {1, 3}},
			want:  []int{1},
		},
		// Add more test cases here...
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinHeightTrees(tt.n, tt.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMinHeightTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
