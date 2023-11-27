package leetcode

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		want  [][]byte
	}{
		{
			name: "case 1",
			board: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			want: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			name: "case 2",
			board: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			want: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			name: "case 3",
			board: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'X', 'X'},
			},
			want: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.board)
			if !reflect.DeepEqual(tt.board, tt.want) {
				t.Errorf("solve() = %v, want %v", tt.board, tt.want)
			}
		})
	}
}
