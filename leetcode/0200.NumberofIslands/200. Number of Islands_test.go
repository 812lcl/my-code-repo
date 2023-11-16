package leetcode

import "testing"

func Test_numIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "test 1",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		{
			name: "test 2",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
		{
			name: "test 3",
			grid: [][]byte{
				{'1', '0', '1', '1', '0', '1', '1'},
				{'1', '0', '1', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '1', '0', '0'},
				{'1', '1', '1', '0', '0', '0', '1'},
				{'1', '1', '0', '1', '1', '0', '1'},
				{'1', '1', '0', '0', '1', '0', '1'},
				{'0', '1', '1', '1', '1', '1', '1'},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands(tt.grid); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
