package leetcode

import "testing"

func Test_arrayNesting(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{5, 4, 0, 3, 1, 6, 2},
			want: 4,
		},
		{
			name: "Example 2",
			nums: []int{0, 1, 2},
			want: 1,
		},
		{
			name: "Example 3",
			nums: []int{0, 2, 1},
			want: 2,
		},
		{
			name: "Empty array",
			nums: []int{},
			want: 0,
		},
		{
			name: "Single element",
			nums: []int{0},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayNesting(tt.nums); got != tt.want {
				t.Errorf("arrayNesting() = %v, want %v", got, tt.want)
			}
			if got := arrayNesting1(tt.nums); got != tt.want {
				t.Errorf("arrayNesting() = %v, want %v", got, tt.want)
			}
		})
	}
}
