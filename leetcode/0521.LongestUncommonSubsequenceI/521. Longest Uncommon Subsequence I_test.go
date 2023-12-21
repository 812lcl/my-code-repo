package leetcode

import "testing"

func Test_findLUSlength(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want int
	}{
		{
			name: "case 1",
			a:    "aba",
			b:    "cdc",
			want: 3,
		},
		{
			name: "case 2",
			a:    "abc",
			b:    "abc",
			want: -1,
		},
		{
			name: "case 3",
			a:    "abcd",
			b:    "efgh",
			want: 4,
		},
		{
			name: "case 4",
			a:    "aaa",
			b:    "bbb",
			want: 3,
		},
		{
			name: "case 5",
			a:    "abc",
			b:    "bca",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLUSlength(tt.a, tt.b); got != tt.want {
				t.Errorf("findLUSlength() = %v, want %v", got, tt.want)
			}
			if got := findLUSlength1(tt.a, tt.b); got != tt.want {
				t.Errorf("findLUSlength() = %v, want %v", got, tt.want)
			}
		})
	}
}
