package leetcode

import "testing"

func Test_minDistance(t *testing.T) {
	tests := []struct {
		name   string
		word1  string
		word2  string
		expect int
	}{
		{
			name:   "empty strings",
			word1:  "",
			word2:  "",
			expect: 0,
		},
		{
			name:   "same word",
			word1:  "hello",
			word2:  "hello",
			expect: 0,
		},
		{
			name:   "one empty word",
			word1:  "hello",
			word2:  "",
			expect: 5,
		},
		{
			name:   "one word is a prefix of the other",
			word1:  "hello",
			word2:  "hell",
			expect: 1,
		},
		{
			name:   "one word is a suffix of the other",
			word1:  "hello",
			word2:  "lo",
			expect: 3,
		},
		{
			name:   "different words",
			word1:  "kitten",
			word2:  "sitting",
			expect: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minDistance(tt.word1, tt.word2)
			if got != tt.expect {
				t.Errorf("minDistance() = %v, want %v", got, tt.expect)
			}
		})
	}
}
