package leetcode

import "testing"

func Test_checkSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{
			name: "empty array",
			nums: []int{},
			k:    3,
			want: false,
		},
		{
			name: "no subarray sum",
			nums: []int{1, 2, 3, 4},
			k:    5,
			want: true,
		},
		{
			name: "subarray sum equals k",
			nums: []int{1, 2, 3, 4},
			k:    6,
			want: true,
		},
		{
			name: "subarray sum divisible by k",
			nums: []int{1, 2, 3, 4},
			k:    3,
			want: true,
		},
		{
			name: "subarray sum not divisible by k",
			nums: []int{1, 2, 3, 4},
			k:    4,
			want: false,
		},
		{
			name: "",
			nums: []int{2, 4, 3},
			k:    6,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSubarraySum(tt.nums, tt.k); got != tt.want {
				t.Errorf("checkSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
