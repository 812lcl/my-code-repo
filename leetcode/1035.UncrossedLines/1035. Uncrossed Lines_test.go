package leetcode

import "testing"

func Test_maxUncrossedLines(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  int
	}{
		{
			name:  "example 1",
			nums1: []int{1, 4, 2},
			nums2: []int{1, 2, 4},
			want:  2,
		},
		{
			name:  "example 2",
			nums1: []int{2, 5, 1, 2, 5},
			nums2: []int{10, 5, 2, 1, 5, 2},
			want:  3,
		},
		{
			name:  "example 3",
			nums1: []int{1, 3, 7, 1, 7, 5},
			nums2: []int{1, 9, 2, 5, 1},
			want:  2,
		},
		{
			name:  "empty arrays",
			nums1: []int{},
			nums2: []int{},
			want:  0,
		},
		{
			name:  "one empty array",
			nums1: []int{1, 2, 3},
			nums2: []int{},
			want:  0,
		},
		{
			name:  "no common elements",
			nums1: []int{1, 2, 3},
			nums2: []int{4, 5, 6},
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxUncrossedLines(tt.nums1, tt.nums2); got != tt.want {
				t.Errorf("maxUncrossedLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
