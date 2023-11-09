package structure

import "fmt"

type color bool

const (
	red   color = true
	black color = false
)

type RBNode struct {
	Value       int
	color       color
	left, right *RBNode
}

type RBTree struct {
	Root *RBNode
}

func (t *RBTree) Insert(value int) {
	t.Root = insert(t.Root, value)
	t.Root.color = black
}

func insert(node *RBNode, value int) *RBNode {
	if node == nil {
		return &RBNode{Value: value, color: red}
	}

	if value < node.Value {
		node.left = insert(node.left, value)
	} else if value > node.Value {
		node.right = insert(node.right, value)
	} else {
		return node
	}

	if isRed(node.right) && !isRed(node.left) {
		node = rotateLeft(node)
	}
	if isRed(node.left) && isRed(node.left.left) {
		node = rotateRight(node)
	}
	if isRed(node.left) && isRed(node.right) {
		flipColors(node)
	}

	return node
}

func isRed(node *RBNode) bool {
	if node == nil {
		return false
	}
	return node.color == red
}

func rotateLeft(node *RBNode) *RBNode {
	right := node.right
	node.right = right.left
	right.left = node
	right.color = node.color
	node.color = red
	return right
}

func rotateRight(node *RBNode) *RBNode {
	left := node.left
	node.left = left.right
	left.right = node
	left.color = node.color
	node.color = red
	return left
}

func flipColors(node *RBNode) {
	node.color = !node.color
	node.left.color = !node.left.color
	node.right.color = !node.right.color
}

func (t *RBTree) Search(value int) bool {
	node := t.Root
	for node != nil {
		if value < node.Value {
			node = node.left
		} else if value > node.Value {
			node = node.right
		} else {
			return true
		}
	}
	return false
}

func (t *RBTree) InOrderTraversal() {
	inOrderTraversal(t.Root)
}

func inOrderTraversal(node *RBNode) {
	if node == nil {
		return
	}
	inOrderTraversal(node.left)
	fmt.Println(node.Value)
	inOrderTraversal(node.right)
}

func RBTreeFunc() {
	root := &RBNode{Value: 1}
	tree := &RBTree{Root: root}
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	tree.InOrderTraversal()
	fmt.Println(tree.Search(3))
	fmt.Println(tree.Search(6))
}
