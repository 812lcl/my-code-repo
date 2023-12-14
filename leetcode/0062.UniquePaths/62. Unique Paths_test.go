package leetcode

import "testing"

func Test_uniquePaths(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		want int
	}{
		{
			name: "case 1",
			m:    3,
			n:    2,
			want: 3,
		},
		{
			name: "case 2",
			m:    7,
			n:    3,
			want: 28,
		},
		{
			name: "case 3",
			m:    1,
			n:    1,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniquePaths(tt.m, tt.n)
			if got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
