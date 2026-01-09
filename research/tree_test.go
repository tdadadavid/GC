package research

import (
	"fmt"
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
	expectedOrder := []int{1, 4, 6, 7, 8, 9, 10}

	result := make([]int, 0)
	result = tree.InOrder(tree.Root, result)
	if len(expectedOrder) != len(result) {
		t.Fatalf("expected length %d, got %d", len(expectedOrder), len(result))
	}

	if isSame := compare(t, expectedOrder, result); !isSame {
		t.Fatalf("expected=%v, got=%v", expectedOrder, result)
	}
}

func TestTree_PostOrder(t *testing.T) {
	tree := NewTree[int]()
	setup(t, tree, 7, 4, 9, 1, 6, 8, 10)
	expectedOrder := []int{1, 6, 4, 8, 10, 9, 7}

	result := make([]int, 0)
	result = tree.PostOrder(tree.Root, result)
	if len(expectedOrder) != len(result) {
		t.Fatalf("expected length %d, got %d", len(expectedOrder), len(result))
	}

	if isSame := compare(t, expectedOrder, result); !isSame {
		t.Fatalf("expected=%v, got=%v", expectedOrder, result)
	}
}

func TestTree_Height(t *testing.T) {
	tests := []struct{
		name string
		input []int
		expectedHeight int
	}{
		{
			name: "Height of an empty tree is -1",
			input: []int{},
			expectedHeight: -1,
		},
		{
			name: "Height of the tree should be 1",
			input: []int{1,2},
			expectedHeight: 1,
		},
		{
			name: "Height of tree should be 4",
			input: []int{10,8,15,3,9,1,2,12},
			expectedHeight: 4,
		},
	}

	for _ , tt := range tests {
		t.Run(tt.name, func (t *testing.T) {
			tree := NewTree[int]()
			for _, value := range tt.input {
				tree.Insert(value)
			}

			count := tree.Height()
			if count != tt.expectedHeight {
				t.Fatalf("expected height of the tree to be %d, got=%d", tt.expectedHeight, count)
			}
		})
	}
}

func TestTree_Minimum(t *testing.T) {
	tests := []struct{
		name string
		input []int
		expectedMin int
		expectError bool
	}{
		{
			name: "the minimum of an empty tree is 0",
			input: []int{},
			expectError: true,
		},
		// {
		// 	name: "minimum is 1",
		// 	input: []int{10,8,15,3,9,1,2,12},
		// 	expectedMin: 1,
		// 	expectError: false,
		// },
	}

	for _ , tt := range tests {
		t.Run(tt.name, func (t *testing.T) {
			tree := NewTree[int]()

			for _, value := range tt.input {
				tree.Insert(value)
			}

			fmt.Printf("tree %v", tree)

			min, err := tree.Min()
			if tt.expectError && err == nil {
				t.Fatalf("expected error got <nil>")
				return
			}
			if min != tt.expectedMin {
				t.Fatalf("expected minimum value of the tree to be %d, got=%d", tt.expectedMin, min)
			}
		})
	}
}

func compare(t *testing.T, expected, got []int) bool {
	t.Helper()

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
