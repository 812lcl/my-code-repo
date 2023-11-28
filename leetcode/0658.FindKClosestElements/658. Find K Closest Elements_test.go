package leetcode

import (
	"reflect"
	"testing"
)

func Test_findClosestElements(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		k      int
		x      int
		expect []int
	}{
		{
			name:   "case 1",
			arr:    []int{1, 2, 3, 4, 5},
			k:      3,
			x:      3,
			expect: []int{2, 3, 4},
		},
		{
			name:   "case 2",
			arr:    []int{1, 2, 3, 4, 5},
			k:      2,
			x:      6,
			expect: []int{4, 5},
		},
		{
			name:   "case 3",
			arr:    []int{1, 2, 3, 4, 5},
			k:      5,
			x:      0,
			expect: []int{1, 2, 3, 4, 5},
		},
		{
			name:   "case 4",
			arr:    []int{1, 2, 3, 4, 5},
			k:      1,
			x:      2,
			expect: []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findClosestElements(tt.arr, tt.k, tt.x)
			if !reflect.DeepEqual(result, tt.expect) {
				t.Errorf("findClosestElements() = %v, want %v", result, tt.expect)
			}
		})
	}
}
