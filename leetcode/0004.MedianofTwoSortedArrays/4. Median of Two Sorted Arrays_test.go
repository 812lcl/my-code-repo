package leetcode

import (
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name   string
		nums1  []int
		nums2  []int
		expect float64
	}{
		{
			name:   "case 1",
			nums1:  []int{1, 3},
			nums2:  []int{2},
			expect: 2.0,
		},
		{
			name:   "case 2",
			nums1:  []int{1, 2},
			nums2:  []int{3, 4},
			expect: 2.5,
		},
		{
			name:   "case 3",
			nums1:  []int{0, 0},
			nums2:  []int{0, 0},
			expect: 0.0,
		},
		{
			name:   "case 4",
			nums1:  []int{},
			nums2:  []int{1},
			expect: 1.0,
		},
		{
			name:   "case 5",
			nums1:  []int{2},
			nums2:  []int{},
			expect: 2.0,
		},
		{
			name:   "case 6",
			nums1:  []int{1, 2, 3, 4, 5, 6},
			nums2:  []int{7, 8, 9, 10, 11, 12},
			expect: 6.5,
		},
		{
			name:   "case 7",
			nums1:  []int{1, 2, 3, 4, 5, 6},
			nums2:  []int{7, 8, 9, 10, 11},
			expect: 6.0,
		},
		{
			name:   "case 8",
			nums1:  []int{1, 2, 3, 4, 5},
			nums2:  []int{6, 7, 8, 9, 10, 11},
			expect: 6,
		},
		{
			name:   "case 9",
			nums1:  []int{1, 2, 3, 4, 5},
			nums2:  []int{6, 7, 8, 9, 10},
			expect: 5.5,
		},
		{
			name:   "case 10",
			nums1:  []int{1, 2, 3, 4, 5},
			nums2:  []int{6, 7, 8, 9},
			expect: 5.0,
		},
		{
			name:   "case 11",
			nums1:  []int{1, 2, 3, 4, 5},
			nums2:  []int{6, 7, 8},
			expect: 4.5,
		},
		{
			name:   "case 12",
			nums1:  []int{1, 2, 3, 4, 5},
			nums2:  []int{6, 7},
			expect: 4.0,
		},
		{
			name:   "case 13",
			nums1:  []int{1, 2, 3, 4, 5},
			nums2:  []int{6},
			expect: 3.5,
		},
		{
			name:   "case 14",
			nums1:  []int{1, 2, 3, 4},
			nums2:  []int{5, 6, 7, 8},
			expect: 4.5,
		},
		{
			name:   "case 15",
			nums1:  []int{1, 2, 3},
			nums2:  []int{4, 5, 6, 7, 8},
			expect: 4.5,
		},
		{
			name:   "case 16",
			nums1:  []int{1, 2},
			nums2:  []int{3, 4, 5, 6, 7, 8},
			expect: 4.5,
		},
		{
			name:   "case 17",
			nums1:  []int{1},
			nums2:  []int{2, 3, 4, 5, 6, 7, 8},
			expect: 4.5,
		},
		{
			name:   "case 18",
			nums1:  []int{1, 2, 3, 4, 5, 6},
			nums2:  []int{1, 2, 3, 4, 5, 6},
			expect: 3.5,
		},
		{
			name:   "case 19",
			nums1:  []int{1, 2, 3, 4, 5, 6},
			nums2:  []int{1, 2, 3, 4, 5},
			expect: 3,
		},
		{
			name:   "case 20",
			nums1:  []int{1, 2, 3, 4, 5},
			nums2:  []int{1, 2, 3, 4, 5, 6},
			expect: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMedianSortedArrays(tt.nums1, tt.nums2)
			if got != tt.expect {
				t.Errorf("findMedianSortedArrays(%v, %v) = %v, but expect %v", tt.nums1, tt.nums2, got, tt.expect)
			}
		})
	}
}
