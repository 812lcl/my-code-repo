package leetcode

import "testing"

func Test_characterReplacement(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			name: "case 0",
			s:    "ABBB",
			k:    2,
			want: 4,
		},
		{
			name: "case 1",
			s:    "ABAB",
			k:    2,
			want: 4,
		},
		{
			name: "case 2",
			s:    "AABABBA",
			k:    1,
			want: 4,
		},
		{
			name: "case 3",
			s:    "AAAA",
			k:    2,
			want: 4,
		},
		{
			name: "case 4",
			s:    "ABCD",
			k:    1,
			want: 2,
		},
		{
			name: "case 5",
			s:    "A",
			k:    0,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := characterReplacement(tt.s, tt.k); got != tt.want {
				t.Errorf("characterReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
