package leetcode

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			name: "example 1",
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "example 2",
			n:    1,
			want: []string{"()"},
		},
		{
			name: "example 3",
			n:    0,
			want: []string{},
		},
		// Add more test cases here...
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := generateParenthesis(tt.n); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			// }
			if got := generateParenthesis1(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
