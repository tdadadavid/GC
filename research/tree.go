package research

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"strings"
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

func (t *Tree[T]) isEmpty() bool {
	return t.Root == nil
}

//--- Helper Functions for test logs ---//

func (t *Tree[T]) String() string {
	if t.isEmpty() {
		return "null"
	}

	return formatNode(t.Root, 0)
}

func formatNode[T constraints.Ordered](n *Node[T], indent int) string {
	if n == nil {
		return "null"
	}

	indentStr := strings.Repeat("  ", indent)
	nextIndentStr := strings.Repeat("  ", indent+1)

	var sb strings.Builder

	sb.WriteString("{\n")

	sb.WriteString(nextIndentStr)
	sb.WriteString(`"value": `)
	sb.WriteString(fmt.Sprintf("%v", n.Value))
	sb.WriteString(",\n")

	sb.WriteString(nextIndentStr)
	sb.WriteString(`"left": `)
	sb.WriteString(formatNode(n.Left, indent+1))
	sb.WriteString(",\n")

	sb.WriteString(nextIndentStr)
	sb.WriteString(`"right": `)
	sb.WriteString(formatNode(n.Right, indent+1))
	sb.WriteString("\n")

	sb.WriteString(indentStr)
	sb.WriteString("}")

	return sb.String()
}
