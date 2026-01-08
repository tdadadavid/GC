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
func (t *Tree[T]) PreOrder() *[]T {
	if t.isEmpty() {
		return nil
	}
}

func (t *Tree[T]) isEmpty() bool {
	return t.Root == nil
}
