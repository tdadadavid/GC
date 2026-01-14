package heaps

import (
	"slices"
	"testing"
)

func TestHeap_Insert(t *testing.T) {
	tests := map[string]struct {
		input         []int
		size          int
		expectedCap   int
		expectedShape []int
	}{
		"Heap should be empty": {
			input:         []int{},
			size:          5,
			expectedCap:   0,
			expectedShape: []int{0, 0, 0, 0, 0},
		},
		"Heap should only have root": {
			input:         []int{1},
			size:          5,
			expectedCap:   1,
			expectedShape: []int{1, 0, 0, 0, 0},
		},
		"Root should be should be 5": {
			input:         []int{1, 4, 5},
			size:          5,
			expectedCap:   3,
			expectedShape: []int{5, 1, 4, 0, 0},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			heap := NewHeap(test.size)

			for _, val := range test.input {
				heap.Insert(val)
			}

			if heap.Size() != test.expectedCap {
				t.Fatalf("expected empty to be %d, got=%d", test.expectedCap, len(heap.values))
			}

			if ok := compare(t, heap.Values(), test.expectedShape); !ok {
				t.Fatalf("expected heap shape to be %d, got=%d", test.expectedShape, heap.Values())
			}
		})
	}
}

// func TestHeap_Remove(t *testing.T) {
// 	tests := map[string]struct {
// 		inputs         []int
// 		valuesToRemove []int
// 		size           int
// 		expectedCap    int
// 		expectedShape  []int
// 	}{
// 		"Root should be z after deletion": {
// 			inputs:         []int{1, 4, 5},
// 			valuesToRemove: []int{5}, //remove root
// 			size:           5,
// 			expectedCap:    2,
// 			expectedShape:  []int{4, 1, 0, 0, 0},
// 		},
// 	}

// 	for name, test := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			t.Parallel()

// 			heap := NewHeap(test.size)

// 			for _, val := range test.inputs {
// 				heap.Insert(val)
// 			}

// 			for _, val := range test.valuesToRemove {
// 				heap.Remove(val)
// 			}

// 			if heap.Size() != test.expectedCap {
// 				t.Fatalf("expected empty to be %d, got=%d", test.expectedCap, len(heap.values))
// 			}

// 			if ok := compare(t, heap.Values(), test.expectedShape); !ok {
// 				t.Fatalf("expected heap shape to be %d, got=%d", test.expectedShape, heap.Values())
// 			}
// 		})
// 	}
// }

func compare(t *testing.T, got, expected []int) bool {
	t.Helper()

	result := slices.Compare(got, expected)
	return result == 0
}
