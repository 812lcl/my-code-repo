package leetcode

import "testing"

func Test_isMatch(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want bool
	}{
		{
			name: "case 1",
			s:    "aa",
			p:    "a",
			want: false,
		},
		{
			name: "case 2",
			s:    "aa",
			p:    "a*",
			want: true,
		},
		{
			name: "case 3",
			s:    "ab",
			p:    ".*",
			want: true,
		},
		{
			name: "case 4",
			s:    "aab",
			p:    "c*a*b",
			want: true,
		},
		{
			name: "case 5",
			s:    "mississippi",
			p:    "mis*is*p*.",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.s, tt.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
