package leetcode

import "testing"

func Test_myAtoi(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "case 1",
			s:    "42",
			want: 42,
		},
		{
			name: "case 2",
			s:    "   -42",
			want: -42,
		},
		{
			name: "case 3",
			s:    "4193 with words",
			want: 4193,
		},
		{
			name: "case 4",
			s:    "words and 987",
			want: 0,
		},
		{
			name: "case 5",
			s:    "-91283472332",
			want: -2147483648,
		},
		{
			name: "case 6",
			s:    "+1",
			want: 1,
		},
		{
			name: "case 7",
			s:    "+-12",
			want: 0,
		},
		{
			name: "case 8",
			s:    "20000000000000000000",
			want: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
