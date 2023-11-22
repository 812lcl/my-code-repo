package leetcode

import (
	"reflect"
	"testing"
)

func Test_medianSlidingWindow(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []float64
	}{
		{
			name: "case 1",
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
			want: []float64{1, -1, -1, 3, 5, 6},
		},
		{
			name: "case 2",
			nums: []int{1, 2, 3, 4, 5},
			k:    1,
			want: []float64{1, 2, 3, 4, 5},
		},
		{
			name: "case 3",
			nums: []int{7, 6, 5, 4, 3, 2, 1},
			k:    4,
			want: []float64{5.5, 4.5, 3.5, 2.5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := medianSlidingWindow(tt.nums, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("medianSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
