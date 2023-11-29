package leetcode

import (
	"fmt"
	"testing"
)

func Test_Trie_StartsWith(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	trie.Insert("banana")
	trie.Insert("orange")

	tests := []struct {
		prefix string
		want   bool
	}{
		{
			prefix: "app",
			want:   true,
		},
		{
			prefix: "ban",
			want:   true,
		},
		{
			prefix: "ora",
			want:   true,
		},
		{
			prefix: "grape",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Prefix: %s", tt.prefix), func(t *testing.T) {
			got := trie.StartsWith(tt.prefix)
			if got != tt.want {
				t.Errorf("StartsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Trie_Search(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	trie.Insert("banana")
	trie.Insert("orange")

	tests := []struct {
		word string
		want bool
	}{
		{
			word: "apple",
			want: true,
		},
		{
			word: "banana",
			want: true,
		},
		{
			word: "orange",
			want: true,
		},
		{
			word: "grape",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Word: %s", tt.word), func(t *testing.T) {
			got := trie.Search(tt.word)
			if got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
