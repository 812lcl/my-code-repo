package leetcode

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "case 2",
			args: args{
				nums: []int{0, 0, 0},
			},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "case 3",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
			if got := threeSum2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_threeSum3(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			want: [][]int{{-1, 0, 1}, {-1, -1, 2}},
		},
		{
			name: "case 2",
			args: args{
				nums: []int{0, 0, 0},
			},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "case 3",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum3(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum3() = %v, want %v", got, tt.want)
			}
		})
	}
}
