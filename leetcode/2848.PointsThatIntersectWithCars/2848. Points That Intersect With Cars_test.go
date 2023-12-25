package leetcode

import "testing"

func Test_numberOfPoints(t *testing.T) {
	tests := []struct {
		name string
		nums [][]int
		want int
	}{
		{
			name: "single interval",
			nums: [][]int{{1, 5}},
			want: 5,
		},
		{
			name: "multiple intervals",
			nums: [][]int{{1, 3}, {5, 7}, {10, 12}},
			want: 9,
		},
		{
			name: "overlapping intervals",
			nums: [][]int{{1, 5}, {3, 7}, {6, 10}},
			want: 10,
		},
		{
			name: "empty intervals",
			nums: [][]int{},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfPoints(tt.nums); got != tt.want {
				t.Errorf("numberOfPoints() = %v, want %v", got, tt.want)
			}
			if got := numberOfPoints1(tt.nums); got != tt.want {
				t.Errorf("numberOfPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
