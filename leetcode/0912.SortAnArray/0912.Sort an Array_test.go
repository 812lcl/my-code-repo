package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "case 1",
			nums: []int{5, 2, 3, 1},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "case 2",
			nums: []int{5, 1, 1, 2, 0, 0},
			want: []int{0, 0, 1, 1, 2, 5},
		},
		{
			name: "case 3",
			nums: []int{1},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArray(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
