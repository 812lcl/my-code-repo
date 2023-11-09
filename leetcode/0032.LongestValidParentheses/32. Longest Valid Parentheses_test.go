package leetcode

import "testing"

func Test_longestValidParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "empty string",
			s:    "",
			want: 0,
		},
		{
			name: "single valid pair",
			s:    "()",
			want: 2,
		},
		{
			name: "single invalid pair",
			s:    ")(",
			want: 0,
		},
		{
			name: "multiple valid pairs",
			s:    "()()()",
			want: 6,
		},
		{
			name: "multiple invalid pairs",
			s:    ")))(((",
			want: 0,
		},
		{
			name: "mixed pairs",
			s:    "(()))(",
			want: 4,
		},
		{
			name: "mixed pairs 2",
			s:    "()(()",
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
			if got := longestValidParentheses1(tt.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
