package leetcode

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expected := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}

	rotate(matrix)

	if !reflect.DeepEqual(matrix, expected) {
		t.Errorf("rotate() = %v, want %v", matrix, expected)
	}
}
