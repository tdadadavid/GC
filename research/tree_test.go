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

	root, left, right := 10, 5, 15
	setup(t, tree, root, left, right)

	if tree.Root.Value != root {
		t.Fatalf("expected value to be 10 got=%d", tree.Root.Value)
	}
	if tree.Root.Left.Value == left && tree.Root.Right.Value != right {
		t.Fatalf("expected right & left subtree to be nil, got left=%d, right=%d", tree.Root.Left.Value, tree.Root.Right.Value)
	}
}

func TestTree_Find(t *testing.T) {
	tree := NewTree[int]()
	setup(t, tree, 10, -20, 30, -40, 50)

	found := tree.Find(-20)

	if !found {
		t.Fatalf("expected found to be true got=%t", found)
	}

	found = tree.Find(1000) // non-existing number
	if found {
		t.Fatalf("expected found to be false got=%t", found)
	}
}

func TestTree_PreOrder(t *testing.T) {
	tree := NewTree[int]()
	setup(t, tree, 7, 4, 9, 1, 6, 8, 10)
	expectedOrder := []int{7, 4, 1, 6, 9, 8, 10}

	result := make([]int, 0)
	result = tree.PreOrder(tree.Root, result)
	if len(expectedOrder) != len(result) {
		t.Fatalf("expected length %d, got %d", len(expectedOrder), len(result))
	}

	if isSame := compare(t, expectedOrder, result); !isSame {
		t.Fatalf("expected=%v, got=%v", expectedOrder, result)
	}
}

func TestTree_InOrder(t *testing.T) {
	tree := NewTree[int]()
	setup(t, tree, 7, 4, 9, 1, 6, 8, 10)
	expectedOrder := []int{1,4,6,7,8,9,10}

	result := make([]int, 0)
	result = tree.InOrder(tree.Root, result)
	if len(expectedOrder) != len(result) {
		t.Fatalf("expected length %d, got %d", len(expectedOrder), len(result))
	}

	if isSame := compare(t, expectedOrder, result); !isSame {
		t.Fatalf("expected=%v, got=%v", expectedOrder, result)
	}
}

func compare(t *testing.T, expected, got []int) bool {
	for i := range len(expected) {
		if expected[i] != got[i] {
			return false
		}
	}
	return true
}

func setup(t *testing.T, tree *Tree[int], values ...any) *Tree[int] {
	t.Helper()

	for _, value := range values {
		tree.Insert(value.(int))
	}

	if tree.Root == nil {
		t.Fatalf("expected root to have node value, got=%T", tree.Root)
	}

	return tree
}
