package leetcode

import (
	"reflect"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want []int
	}{
		{
			name: "case 1",
			s:    "cbaebabacd",
			p:    "abc",
			want: []int{0, 6},
		},
		{
			name: "case 2",
			s:    "abab",
			p:    "ab",
			want: []int{0, 1, 2},
		},
		{
			name: "case 3",
			s:    "aaaaaa",
			p:    "aaa",
			want: []int{0, 1, 2, 3},
		},
		{
			name: "case 4",
			s:    "abcdefg",
			p:    "xyz",
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagrams(tt.s, tt.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
