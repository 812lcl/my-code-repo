package leetcode

import "testing"

func Test_evalRPN(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		{
			name:   "example 1",
			tokens: []string{"2", "1", "+", "3", "*"},
			want:   9,
		},
		{
			name:   "example 2",
			tokens: []string{"4", "13", "5", "/", "+"},
			want:   6,
		},
		{
			name:   "example 3",
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:   22,
		},
		{
			name:   "single number",
			tokens: []string{"42"},
			want:   42,
		},
		{
			name:   "division by zero",
			tokens: []string{"1", "0", "/"},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalRPN(tt.tokens); got != tt.want {
				t.Errorf("evalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}
