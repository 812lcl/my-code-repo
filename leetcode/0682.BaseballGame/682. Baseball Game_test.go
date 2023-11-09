package leetcode

import "testing"

func TestCalPoints(t *testing.T) {
	tests := []struct {
		name       string
		operations []string
		want       int
	}{
		{
			name:       "example 1",
			operations: []string{"5", "2", "C", "D", "+"},
			want:       30,
		},
		{
			name:       "example 2",
			operations: []string{"5", "-2", "4", "C", "D", "9", "+", "+"},
			want:       27,
		},
		{
			name:       "empty input",
			operations: []string{},
			want:       0,
		},
		{
			name:       "single input",
			operations: []string{"5"},
			want:       5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calPoints(tt.operations); got != tt.want {
				t.Errorf("calPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
