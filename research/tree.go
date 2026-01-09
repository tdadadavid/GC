package research

import (
	"errors"
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
	if isEmpty(t.Root) {
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
	if isEmpty(t.Root) {
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

func (t *Tree[T]) Min() (T, error) {
	result, ok := minimum[T](t.Root)

	if !ok {
		return result, errors.New("error finding min")
	}

	return result, nil
}

func isEmpty[T constraints.Ordered](node *Node[T]) bool {
	return node == nil
}



func minimum[T constraints.Ordered](root *Node[T]) (T, bool) {
	var zero T

	if isEmpty(root) {
		return zero, false
	}

	if isLeaf(root) {
		return root.Value, false
	}

	left, ok := minimum(root.Left)
	if !ok {
		return left, false
	}

	right, ok := minimum(root.Right)
	if !ok {
		return right, false
	}

	mini := min(left, right)

	return min(mini, root.Value), true
}

func isLeaf[T constraints.Ordered](node *Node[T]) bool {
	return node.Left == nil && node.Right == nil
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
