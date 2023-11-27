package leetcode

import (
	"reflect"
	"testing"
)

func Test_cloneGraph(t *testing.T) {
	// Test case 1
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node1.Neighbors = append(node1.Neighbors, node2, node4)
	node2.Neighbors = append(node2.Neighbors, node1, node3)
	node3.Neighbors = append(node3.Neighbors, node2, node4)
	node4.Neighbors = append(node4.Neighbors, node1, node3)

	want1 := &Node{Val: 1}
	want2 := &Node{Val: 2}
	want3 := &Node{Val: 3}
	want4 := &Node{Val: 4}
	want1.Neighbors = append(want1.Neighbors, want2, want4)
	want2.Neighbors = append(want2.Neighbors, want1, want3)
	want3.Neighbors = append(want3.Neighbors, want2, want4)
	want4.Neighbors = append(want4.Neighbors, want1, want3)

	if got := cloneGraph(node1); !reflect.DeepEqual(got, want1) {
		t.Errorf("cloneGraph() = %v, want %v", got, want1)
	}

	// Test case 2
	node5 := &Node{Val: 5}
	want5 := &Node{Val: 5}

	if got := cloneGraph(node5); !reflect.DeepEqual(got, want5) {
		t.Errorf("cloneGraph() = %v, want %v", got, want5)
	}

	// Test case 3
	var node6 *Node
	var want6 *Node

	if got := cloneGraph(node6); !reflect.DeepEqual(got, want6) {
		t.Errorf("cloneGraph() = %v, want %v", got, want6)
	}
}
