package leetcode

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		s      string
		expect [][]string
	}{
		{
			s:      "aab",
			expect: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			s:      "a",
			expect: [][]string{{"a"}},
		},
		{
			s:      "abc",
			expect: [][]string{{"a", "b", "c"}},
		},
	}

	for _, test := range tests {
		result := partition(test.s)
		if !reflect.DeepEqual(result, test.expect) {
			t.Errorf("Test case failed: %s", test.s)
		}
	}
}
