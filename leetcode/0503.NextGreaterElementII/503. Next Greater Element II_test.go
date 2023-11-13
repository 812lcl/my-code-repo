package leetcode

import (
	"reflect"
	"testing"
)

func Test_nextGreaterElements(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 5",
			args: args{
				nums: []int{100, 1, 11, 1, 120, 111, 123, 1, -1, -100},
			},
			want: []int{120, 11, 120, 120, 123, 123, -1, 100, 100, 100},
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 3, 4, 2},
			},
			want: []int{3, 4, -1, 3},
		},
		{
			name: "case 3",
			args: args{
				nums: []int{5, 4, 3, 2, 1},
			},
			want: []int{-1, 5, 5, 5, 5},
		},
		{
			name: "case 4",
			args: args{
				nums: []int{1},
			},
			want: []int{-1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElements(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElements() = %v, want %v", got, tt.want)
			}
			if got := nextGreaterElements1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
