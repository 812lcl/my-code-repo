package leetcode

import (
	"testing"
)

func Test_connect(t *testing.T) {
	// Test case 1: Single node
	root1 := &Node{Val: 1}
	connect(root1)
	if root1.Next != nil {
		t.Errorf("Test case 1 failed: Expected root1.Next to be nil, got %v", root1.Next)
	}
	connect2(root1)
	if root1.Next != nil {
		t.Errorf("Test case 1 failed: Expected root1.Next to be nil, got %v", root1.Next)
	}

}
