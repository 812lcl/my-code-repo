package leetcode

import "testing"

func TestWordDictionary_Search(t *testing.T) {
	// Create a new WordDictionary
	wordDict := Constructor()

	// Add words to the dictionary
	wordDict.AddWord("apple")
	wordDict.AddWord("banana")
	wordDict.AddWord("cherry")

	// Test cases
	tests := []struct {
		name   string
		word   string
		expect bool
	}{
		{
			name:   "Existing word",
			word:   "apple",
			expect: true,
		},
		{
			name:   "Non-existing word",
			word:   "grape",
			expect: false,
		},
		{
			name:   "Word with wildcard",
			word:   "b.n.n.",
			expect: true,
		},
		{
			name:   "Word with wildcard - no match",
			word:   "b.n.n.a",
			expect: false,
		},
	}

	// Run the tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordDict.Search(tt.word)
			if got != tt.expect {
				t.Errorf("Search(%s) = %v, want %v", tt.word, got, tt.expect)
			}
		})
	}
}
