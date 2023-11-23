package leetcode

import (
	"testing"
)

func Test_reorganizeString(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "Example 3",
			input:  "baaba",
			output: "ababa",
		},
		{
			name:   "Example 1",
			input:  "aab",
			output: "aba",
		},
		{
			name:   "Example 2",
			input:  "aaab",
			output: "",
		},
		{
			name:   "Empty String",
			input:  "",
			output: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorganizeString(tt.input); got != tt.output {
				t.Errorf("reorganizeString() = %v, want %v", got, tt.output)
			}
		})
	}
}
