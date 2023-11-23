package leetcode

import (
	"reflect"
	"testing"
)

func Test_rearrangeBarcodes(t *testing.T) {
	tests := []struct {
		name     string
		barcodes []int
		want     []int
	}{
		{
			name:     "case 1",
			barcodes: []int{1, 1, 1, 2, 2},
			want:     []int{1, 2, 1, 2, 1},
		},
		{
			name:     "case 2",
			barcodes: []int{1, 1, 1, 2, 2, 3},
			want:     []int{1, 2, 1, 2, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rearrangeBarcodes(tt.barcodes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rearrangeBarcodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
