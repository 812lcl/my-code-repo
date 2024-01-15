package leetcode

import "testing"

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		// Testing an empty array
		{[]int{}, 0, 0},

		// Testing a single element array that matches k
		{[]int{5}, 5, 1},

		// Testing a single element array that doesn't match k
		{[]int{5}, 10, 0},

		// Testing an array with multiple elements that matches k
		{[]int{1, 2, 3, 4}, 6, 1},

		// Testing an array with multiple elements that doesn't match k
		{[]int{1, 2, 3, 4}, 10, 1},

		// Testing an array with negative numbers that matches k
		{[]int{-1, -2, -3, -4}, -3, 2},

		// Testing an array with multiple elements that matches k
		{[]int{1, -1, 0}, 0, 3},
	}

	for _, tt := range tests {
		got := subarraySum(tt.nums, tt.k)
		if got != tt.want {
			t.Errorf("subarraySum(%v, %v) = %v, want %v", tt.nums, tt.k, got, tt.want)
		}
		got = subarraySum1(tt.nums, tt.k)
		if got != tt.want {
			t.Errorf("subarraySum(%v, %v) = %v, want %v", tt.nums, tt.k, got, tt.want)
		}
	}
}
