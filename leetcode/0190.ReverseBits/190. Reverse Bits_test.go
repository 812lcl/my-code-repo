package leetcode

import "testing"

func Test_reverseBits(t *testing.T) {
	tests := []struct {
		name string
		num  uint32
		want uint32
	}{
		{
			name: "case 1",
			num:  0b00000000000000000000000000000000,
			want: 0b00000000000000000000000000000000,
		},
		{
			name: "case 2",
			num:  0b11111111111111111111111111111111,
			want: 0b11111111111111111111111111111111,
		},
		{
			name: "case 3",
			num:  0b00000000000000000000000000001111,
			want: 0b11110000000000000000000000000000,
		},
		{
			name: "case 4",
			num:  0b10101010101010101010101010101010,
			want: 0b01010101010101010101010101010101,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits(tt.num); got != tt.want {
				t.Errorf("reverseBits() = %b, want %b", got, tt.want)
			}
		})
	}
}
