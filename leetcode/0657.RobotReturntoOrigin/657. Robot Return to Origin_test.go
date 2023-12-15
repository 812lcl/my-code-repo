package leetcode

import "testing"

func Test_judgeCircle(t *testing.T) {
	tests := []struct {
		name  string
		moves string
		want  bool
	}{
		{
			name:  "moves in a circle",
			moves: "UDLR",
			want:  true,
		},
		{
			name:  "moves in a straight line",
			moves: "UUUU",
			want:  false,
		},
		{
			name:  "moves in a zigzag pattern",
			moves: "UDUDUD",
			want:  true,
		},
		{
			name:  "no moves",
			moves: "",
			want:  true,
		},
		{
			name:  "moves in a square",
			moves: "UDLRUDLR",
			want:  true,
		},
		{
			name:  "moves in a diagonal",
			moves: "URDL",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgeCircle(tt.moves); got != tt.want {
				t.Errorf("judgeCircle() = %v, want %v", got, tt.want)
			}
		})
	}
}
