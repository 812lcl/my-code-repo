package leetcode

import "testing"

func Test_rangeBitwiseAnd(t *testing.T) {
	tests := []struct {
		name  string
		left  int
		right int
		want  int
	}{
		{
			name:  "case 1",
			left:  5,
			right: 7,
			want:  4,
		},
		{
			name:  "case 2",
			left:  0,
			right: 1,
			want:  0,
		},
		{
			name:  "case 3",
			left:  1,
			right: 1,
			want:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeBitwiseAnd(tt.left, tt.right); got != tt.want {
				t.Errorf("rangeBitwiseAnd() = %v, want %v", got, tt.want)
			}
			if got := rangeBitwiseAnd1(tt.left, tt.right); got != tt.want {
				t.Errorf("rangeBitwiseAnd() = %v, want %v", got, tt.want)
			}
			if got := rangeBitwiseAnd2(tt.left, tt.right); got != tt.want {
				t.Errorf("rangeBitwiseAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
