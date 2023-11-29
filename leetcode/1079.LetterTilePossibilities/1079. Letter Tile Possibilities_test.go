package leetcode

import "testing"

func Test_numTilePossibilities1(t *testing.T) {
	tests := []struct {
		name   string
		tiles  string
		result int
	}{
		{
			name:   "case 1",
			tiles:  "AAB",
			result: 8,
		},
		{
			name:   "case 2",
			tiles:  "AAABBC",
			result: 188,
		},
		{
			name:   "case 3",
			tiles:  "v",
			result: 1,
		},
		{
			name:   "case 4",
			tiles:  "XYZ",
			result: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTilePossibilities(tt.tiles); got != tt.result {
				t.Errorf("numTilePossibilities1() = %v, want %v", got, tt.result)
			}
			if got := numTilePossibilities1(tt.tiles); got != tt.result {
				t.Errorf("numTilePossibilities1() = %v, want %v", got, tt.result)
			}
		})
	}
}
