package leetcode

import "testing"

func TestTopK(t *testing.T) {
	tests := []struct {
		name string
		k    int
		nums []int
		want int
	}{
		{
			name: "case 4",
			k:    7,
			nums: []int{4, 5, 8, 2, 3, 5, 10, 9, 4},
			want: 4,
		},
		{
			name: "case 1",
			k:    3,
			nums: []int{4, 5, 1, 6, 2, 7, 3, 8},
			want: 6,
		},
		{
			name: "case 2",
			k:    2,
			nums: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			want: 8,
		},
		{
			name: "case 3",
			k:    1,
			nums: []int{1, 2, 3, 4, 5},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TopK(tt.k, tt.nums); got != tt.want {
				t.Errorf("TopK() = %v, want %v", got, tt.want)
			}
			if got := QuickTopK(tt.k, tt.nums); got != tt.want {
				t.Errorf("TopK() = %v, want %v", got, tt.want)
			}
		})
	}
}
