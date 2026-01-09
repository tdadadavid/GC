package research

import (
	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

func NewNode[T constraints.Ordered](value T) (n *Node[T]) {
	n = &Node[T]{
		Value: value,
	}
	return n
}

type Tree[T constraints.Ordered] struct {
	Root *Node[T]
}

func NewTree[T constraints.Ordered]() *Tree[T] {
	return &Tree[T]{}
}

func (t *Tree[T]) Insert(value T) {
	node := NewNode(value)

	// first check if the root of the tree is empty
	if t.isEmpty() {
		t.Root = node
		return
	}

	currentNode := t.Root
	for {
		// 	a. less than root -> Left
		if value <= currentNode.Value {
			// check if the node is empty
			if currentNode.Left == nil {
				currentNode.Left = node
				break
			}
			// move to the currentNode
			currentNode = currentNode.Left
		} else { //b. greater than root  -> Right
			// check if the node is empty
			if currentNode.Right == nil {
				currentNode.Right = node
				break
			}
			currentNode = currentNode.Right
		}

	}

}

func (t *Tree[T]) Find(value T) (isFound bool) {
	// if the tree is empty return false
	if t.isEmpty() {
		return false
	}

	currentNode := t.Root
	for {
		if value < currentNode.Value {
			if currentNode.Left == nil {
				return false
			}
			currentNode = currentNode.Left
		} else if value > currentNode.Value {
			if currentNode.Right == nil {
				return false
			}
			currentNode = currentNode.Right
		} else {
			return true
		}
	}
}

// Root -> Left -> Right
func (t *Tree[T]) PreOrder(node *Node[T], result []T) []T {
	if node == nil {
		return result
	}

	// Root
	result = append(result, node.Value)

	// Left
	result = t.PreOrder(node.Left, result)

	// Right
	result = t.PreOrder(node.Right, result)

	return result
}

// Left -> Root -> Right
func (t *Tree[T]) InOrder(node *Node[T], result []T) []T {
	if node == nil {
		return result
	}

	result = t.InOrder(node.Left, result)
	result = append(result, node.Value)
	result = t.InOrder(node.Right, result)

	return result
}

// Left -> Right -> Root
func (t *Tree[T]) PostOrder(node *Node[T], result []T) []T {
	if node == nil {
		return result
	}

	result = t.PostOrder(node.Left, result)
	result = t.PostOrder(node.Right, result)
	result = append(result, node.Value)

	return result
}

// Height of tree is the number of edges from the leaf node to the particular node.
func (t *Tree[T]) Height() (count int) {
	return height[T](t.Root)
}

func (t *Tree[T]) isEmpty() bool {
	return t.Root == nil
}

func height[T constraints.Ordered](root *Node[T]) (count int) {
	if root == nil { // this is an empty tree or a leaf node
		return -1 // the height of an empty is -1
	}

	if root.Left == nil && root.Right == nil {
		return count // count is zero at this point
	}

	count = 1 + max(height(root.Left), height(root.Right))

	return count
}
