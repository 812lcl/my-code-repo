package leetcode

import "testing"

func Test_removeOuterParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "no outer parentheses",
			s:    "abc",
			want: "abc",
		},
		{
			name: "one pair of outer parentheses",
			s:    "(abc)",
			want: "abc",
		},
		{
			name: "multiple pairs of outer parentheses",
			s:    "(()())(())",
			want: "()()()",
		},
		{
			name: "nested parentheses",
			s:    "(()())(()(()))",
			want: "()()()(())",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeOuterParentheses(tt.s); got != tt.want {
				t.Errorf("removeOuterParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
