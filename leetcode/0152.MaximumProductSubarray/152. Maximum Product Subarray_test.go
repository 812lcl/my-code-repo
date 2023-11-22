package leetcode

import "testing"

// Test cases
func Test_maxProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "case 1",
			nums: []int{2, 3, -2, 4},
			want: 6,
		},
		{
			name: "case 2",
			nums: []int{-2, 0, -1},
			want: 0,
		},
		{
			name: "case 3",
			nums: []int{-2, 3, -4},
			want: 24,
		},
		{
			name: "case 4",
			nums: []int{0, 2},
			want: 2,
		},
		{
			name: "case 5",
			nums: []int{-2},
			want: -2,
		},
		{
			name: "case 6",
			nums: []int{0},
			want: 0,
		},
		{
			name: "case 7",
			nums: []int{-2, -3, -4},
			want: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
