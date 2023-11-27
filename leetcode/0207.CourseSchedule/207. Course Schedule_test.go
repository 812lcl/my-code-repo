package leetcode

import "testing"

func Test_canFinish(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{
			name:          "example 0",
			numCourses:    4,
			prerequisites: [][]int{{1, 2}, {2, 3}},
			want:          true,
		},
		{
			name:          "example 1",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			want:          true,
		},
		{
			name:          "example 2",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			want:          false,
		},
		{
			name:          "example 3",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}},
			want:          true,
		},
		{
			name:          "example 4",
			numCourses:    3,
			prerequisites: [][]int{{1, 0}, {2, 1}, {0, 2}},
			want:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.numCourses, tt.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
			if got := canFinish1(tt.numCourses, tt.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
