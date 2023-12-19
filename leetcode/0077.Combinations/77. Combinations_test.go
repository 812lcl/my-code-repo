package leetcode

import (
	"reflect"
	"testing"
)

func Test_combine(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want [][]int
	}{
		{
			name: "case 1",
			n:    4,
			k:    2,
			want: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			name: "case 2",
			n:    3,
			k:    3,
			want: [][]int{{1, 2, 3}},
		},
		{
			name: "case 3",
			n:    5,
			k:    1,
			want: [][]int{{1}, {2}, {3}, {4}, {5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combine(tt.n, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}
