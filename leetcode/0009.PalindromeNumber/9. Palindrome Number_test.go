package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{
			name: "positive palindrome",
			x:    121,
			want: true,
		},
		{
			name: "negative palindrome",
			x:    -121,
			want: false,
		},
		{
			name: "non-palindrome",
			x:    123,
			want: false,
		},
		{
			name: "single digit",
			x:    5,
			want: true,
		},
		{
			name: "zero",
			x:    0,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
			if got := isPalindrome1(tt.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
