package leetcode

import (
	"reflect"
	"testing"
)

func Test_findThePrefixCommonArray(t *testing.T) {
	tests := []struct {
		name string
		A    []int
		B    []int
		want []int
	}{
		{
			name: "case 1",
			A:    []int{1, 2, 3, 4, 5},
			B:    []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "case 2",
			A:    []int{1, 2, 3, 4, 5},
			B:    []int{5, 4, 3, 2, 1},
			want: []int{0, 0, 1, 3, 5},
		},
		{
			name: "case 3",
			A:    []int{1, 2, 3, 4, 5},
			B:    []int{6, 7, 8, 9, 10},
			want: []int{0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findThePrefixCommonArray(tt.A, tt.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findThePrefixCommonArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
