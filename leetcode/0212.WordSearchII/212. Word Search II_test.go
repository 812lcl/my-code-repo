package leetcode

import (
	"reflect"
	"testing"
)

func Test_findWords(t *testing.T) {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	want := []string{"oath", "eat"}

	got := findWords(board, words)
	got1 := findWords1(board, words)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("findWords() = %v, want %v", got, want)
	}
	if !reflect.DeepEqual(got1, want) {
		t.Errorf("findWords() = %v, want %v", got, want)
	}
}
