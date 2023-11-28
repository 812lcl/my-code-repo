package leetcode

import (
	"reflect"
	"testing"
)

func Test_findRepeatedDnaSequences(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{
			name: "test case 1",
			s:    "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",
			want: []string{"AAAAACCCCC", "CCCCCAAAAA"},
		},
		{
			name: "test case 2",
			s:    "AAAAAAAAAAA",
			want: []string{"AAAAAAAAAA"},
		},
		{
			name: "test case 3",
			s:    "ACGTACGTACGT",
			want: []string{},
		},
		// Add more test cases here...
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatedDnaSequences(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRepeatedDnaSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
