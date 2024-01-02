package leetcode

import "testing"

func Test_minOperations(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "case 1",
			n:    5,
			want: 2,
		},
		{
			name: "case 2",
			n:    8,
			want: 1,
		},
		{
			name: "case 3",
			n:    10,
			want: 2,
		},
		{
			name: "case 4",
			n:    15,
			want: 2,
		},
		{
			name: "case 5",
			n:    1,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.n); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
