package leetcode

import (
	"reflect"
	"testing"
)

func Test_findOrder(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          []int
	}{
		{
			name:          "case 1",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			want:          []int{0, 1, 2, 3},
		},
		{
			name:          "case 2",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			want:          []int{0, 1},
		},
		{
			name:          "case 3",
			numCourses:    3,
			prerequisites: [][]int{{1, 0}, {2, 1}},
			want:          []int{0, 1, 2},
		},
		{
			name:          "case 4",
			numCourses:    3,
			prerequisites: [][]int{{1, 0}, {0, 2}, {2, 1}},
			want:          []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findOrder(tt.numCourses, tt.prerequisites)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
