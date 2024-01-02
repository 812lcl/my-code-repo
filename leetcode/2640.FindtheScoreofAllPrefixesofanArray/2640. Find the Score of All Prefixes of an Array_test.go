package leetcode

import (
	"reflect"
	"testing"
)

func Test_findPrefixScore(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int64
	}{
		{
			name: "case 1",
			nums: []int{1, 2, 3, 4, 5},
			want: []int64{2, 6, 12, 20, 30},
		},
		{
			name: "case 2",
			nums: []int{5, 4, 3, 2, 1},
			want: []int64{10, 19, 27, 34, 40},
		},
		{
			name: "case 3",
			nums: []int{1, 1, 1, 1, 1},
			want: []int64{2, 4, 6, 8, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPrefixScore(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPrefixScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
