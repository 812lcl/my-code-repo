package leetcode

import "testing"

func Test_reverseWords(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "case 1",
			s:    "the sky is blue",
			want: "blue is sky the",
		},
		{
			name: "case 2",
			s:    "  hello world!  ",
			want: "world! hello",
		},
		{
			name: "case 3",
			s:    "a good   example",
			want: "example good a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
			if got := reverseWords1(tt.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
