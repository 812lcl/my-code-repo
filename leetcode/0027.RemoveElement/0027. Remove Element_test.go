package leetcode

import (
	"reflect"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want: 5,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{1, 2, 3},
				val:  4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
			if got := tt.args.nums[:tt.want]; !reflect.DeepEqual(got, tt.args.nums[:tt.want]) {
				t.Errorf("removeElement() = %v, want %v", got, tt.args.nums[:tt.want])
			}
		})
	}
}
