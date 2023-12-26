package leetcode

import "testing"

func Test_longestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name  string
		text1 string
		text2 string
		want  int
	}{
		{
			name:  "case 1",
			text1: "abcde",
			text2: "ace",
			want:  3,
		},
		{
			name:  "case 2",
			text1: "abc",
			text2: "def",
			want:  0,
		},
		{
			name:  "case 3",
			text1: "abc",
			text2: "abc",
			want:  3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.text1, tt.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
