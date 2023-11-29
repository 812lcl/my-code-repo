package leetcode

import "testing"

func Test_strStr(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		want     int
	}{
		{
			name:     "Example 1",
			haystack: "hello",
			needle:   "ll",
			want:     2,
		},
		{
			name:     "Example 2",
			haystack: "aaaaa",
			needle:   "bba",
			want:     -1,
		},
		{
			name:     "Example 3",
			haystack: "",
			needle:   "",
			want:     0,
		},
		{
			name:     "Empty Needle",
			haystack: "hello",
			needle:   "",
			want:     0,
		},
		{
			name:     "Empty Haystack",
			haystack: "",
			needle:   "ll",
			want:     -1,
		},
		{
			name:     "Needle Longer Than Haystack",
			haystack: "hello",
			needle:   "hello world",
			want:     -1,
		},
		{
			name:     "Needle at the Beginning",
			haystack: "hello",
			needle:   "he",
			want:     0,
		},
		{
			name:     "Needle at the End",
			haystack: "hello",
			needle:   "lo",
			want:     3,
		},
		{
			name:     "Multiple Occurrences",
			haystack: "hello world",
			needle:   "o",
			want:     4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.haystack, tt.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
