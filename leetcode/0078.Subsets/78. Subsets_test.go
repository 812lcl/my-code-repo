package leetcode

import (
	"reflect"
	"testing"
)

func Test_subsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3},
			want: [][]int{
				{},
				{1},
				{2},
				{3},
				{1, 2},
				{1, 3},
				{2, 3},
				{1, 2, 3},
			},
		},
		{
			name: "Example 2",
			nums: []int{0},
			want: [][]int{
				{},
				{0},
			},
		},
		{
			name: "Example 3",
			nums: []int{1, 2},
			want: [][]int{
				{},
				{1},
				{2},
				{1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subsets1(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3},
			want: [][]int{
				{},
				{1},
				{2},
				{1, 2},
				{3},
				{1, 3},
				{2, 3},
				{1, 2, 3},
			},
		},
		{
			name: "Example 2",
			nums: []int{0},
			want: [][]int{
				{},
				{0},
			},
		},
		{
			name: "Example 3",
			nums: []int{1, 2},
			want: [][]int{
				{},
				{1},
				{2},
				{1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets1(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
			if got := subsets2(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
