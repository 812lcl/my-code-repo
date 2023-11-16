package leetcode

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func Test_generateTrees(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []*TreeNode
	}{
		{
			name: "n=0",
			n:    0,
			want: nil,
		},
		{
			name: "n=1",
			n:    1,
			want: []*TreeNode{
				{Val: 1},
			},
		},
		{
			name: "n=2",
			n:    2,
			want: []*TreeNode{
				{Val: 1, Right: &TreeNode{Val: 2}},
				{Val: 2, Left: &TreeNode{Val: 1}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateTrees(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateTrees() = %v, want %v", spew.Sdump(got), spew.Sdump(tt.want))
			}
		})
	}
}
