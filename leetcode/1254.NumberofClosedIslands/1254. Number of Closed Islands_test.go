package leetcode

import "testing"

func TestClosedIsland(t *testing.T) {
	grid1 := [][]int{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
	}
	if res := closedIsland(grid1); res != 1 {
		t.Errorf("closedIsland() = %d, want 1", res)
	}

	grid2 := [][]int{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
	}
	if res := closedIsland(grid2); res != 1 {
		t.Errorf("closedIsland() = %d, want 1", res)
	}

	grid3 := [][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}
	if res := closedIsland(grid3); res != 1 {
		t.Errorf("closedIsland() = %d, want 1", res)
	}

	grid4 := [][]int{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}
	if res := closedIsland(grid4); res != 0 {
		t.Errorf("closedIsland() = %d, want 0", res)
	}
}
