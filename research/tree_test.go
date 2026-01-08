package research

import (
	"testing"
)

func TestTree_InsertRoot(t *testing.T) {
	tree := NewTree[int]()

	value := 10

	tree.Insert(value)

	if tree.Root == nil {
		t.Fatalf("expected root to have node value, got=%T", tree.Root)
	}
	if tree.Root.Value != value {
		t.Fatalf("expected value to be 10 got=%d", tree.Root.Value)
	}
	if tree.Root.Left != nil && tree.Root.Right != nil {
		t.Fatalf("expected right & left subtree to be nil, got left=%T, right=%T", tree.Root.Left, tree.Root.Right)
	}

}

func TestTree_Insert(t *testing.T) {
	tree := NewTree[int]()

	root := 10
	left := 5
	right := 15

	tree.Insert(root)
	tree.Insert(left)
	tree.Insert(right)

	if tree.Root == nil {
		t.Fatalf("expected root to have node value, got=%T", tree.Root)
	}
	if tree.Root.Value != root {
		t.Fatalf("expected value to be 10 got=%d", tree.Root.Value)
	}
	if tree.Root.Left.Value == left && tree.Root.Right.Value != right {
		t.Fatalf("expected right & left subtree to be nil, got left=%d, right=%d", tree.Root.Left.Value, tree.Root.Right.Value)
	}
}
