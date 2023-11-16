package leetcode

import "testing"

func Test_minPathSum(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "case 1",
			grid: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			want: 7,
		},
		{
			name: "case 2",
			grid: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			want: 12,
		},
		{
			name: "case 3",
			grid: [][]int{
				{1, 2},
				{3, 4},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum(tt.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
			if got := minPathSum1(tt.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
