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

			if ok := compare(heap.Values(), test.expectedShape); !ok {
				t.Fatalf("expected heap shape to be %d, got=%d", heap.Values(), test.expectedShape)
			}
		})
	}
}

func compare(got, expected []int) bool {
	result := slices.Compare(got, expected)
	return result == 0
}
