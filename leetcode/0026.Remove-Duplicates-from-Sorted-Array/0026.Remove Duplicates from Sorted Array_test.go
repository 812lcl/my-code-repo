package leetcode

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want: 5,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
			if got := tt.args.nums[:tt.want]; !reflect.DeepEqual(got, tt.args.nums[:tt.want]) {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.args.nums[:tt.want])
			}
		})
	}
}
