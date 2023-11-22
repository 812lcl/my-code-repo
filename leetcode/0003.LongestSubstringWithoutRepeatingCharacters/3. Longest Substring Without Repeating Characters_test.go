package leetcode

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output int
	}{
		{
			name:   "Empty string",
			input:  "",
			output: 0,
		},
		{
			name:   "String with no repeating characters",
			input:  "abcde",
			output: 5,
		},
		{
			name:   "String with repeating characters",
			input:  "aabccdef",
			output: 4,
		},
		{
			name:   "String with all repeating characters",
			input:  "aaaaa",
			output: 1,
		},
		{
			name:   "String with one character",
			input:  "a",
			output: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lengthOfLongestSubstring(tt.input)
			if result != tt.output {
				t.Errorf("lengthOfLongestSubstring() = %d, want %d", result, tt.output)
			}
		})
	}
}
