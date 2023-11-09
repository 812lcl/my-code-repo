package leetcode

import (
	"reflect"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	testCases := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{[]int{4, 1, 2}, []int{1, 3, 4, 2}, []int{-1, 3, -1}},
		{[]int{2, 4}, []int{1, 2, 3, 4}, []int{3, -1}},
		{[]int{1, 2, 3}, []int{3, 2, 1}, []int{-1, -1, -1}},
		{[]int{1, 3, 5, 2, 4}, []int{6, 5, 4, 3, 2, 1, 7}, []int{7, 7, 7, 7, 7}},
	}

	for _, tc := range testCases {
		got := nextGreaterElement(tc.nums1, tc.nums2)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("nextGreaterElement(%v, %v) = %v; want %v", tc.nums1, tc.nums2, got, tc.want)
		}
		got = nextGreaterElement1(tc.nums1, tc.nums2)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("nextGreaterElement(%v, %v) = %v; want %v", tc.nums1, tc.nums2, got, tc.want)
		}
	}
}
