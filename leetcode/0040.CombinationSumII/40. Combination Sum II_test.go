package leetcode

import (
	"reflect"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	expected := [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}
	result := combinationSum2(candidates, target)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("combinationSum2() = %v, want %v", result, expected)
	}
}
