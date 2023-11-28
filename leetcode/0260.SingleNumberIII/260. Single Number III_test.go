package leetcode

import (
	"reflect"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "case 1",
			nums: []int{1, 2, 1, 3, 2, 5},
			want: []int{5, 3},
		},
		{
			name: "case 2",
			nums: []int{2, 4, 6, 8, 10, 2, 4, 6},
			want: []int{8, 10},
		},
		{
			name: "case 3",
			nums: []int{1, 1, 2, 2, 3, 4, 4, 5},
			want: []int{5, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
