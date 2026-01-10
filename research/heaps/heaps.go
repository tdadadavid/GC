package heaps

import (
	"golang.org/x/exp/constraints"
)

// Root
//   Left = (2 * n) + 1
//   Right = (2 * n) + 2
//   Parent = (index-1)/2

type Heap struct {
	values []int
	size   int
}

func NewHeap(size int) *Heap {
	return &Heap{
		values: make([]int, size),
		size:   0,
	}
}

func (h *Heap) Insert(value int) {
	// insert at the current position (size)
	h.values[h.size] = value
	h.size++

	h.bubbleUp()
}

func (h *Heap) Remove(value int) {

}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) Values() []int {
	return h.values
}

func (h *Heap) bubbleUp() {
	// get the index of the just inserted value
	index := h.size - 1

	// keep bubbling up until heap is balance
	for index > 0 && h.values[index] > h.parent(index) { // here we implement a MaxHeap because condition is child > parent
		h.swap(index, h.parent(index))

		// bubble up
		index = h.parent(index)
	}
}

func (h *Heap) swap(child, parent int) {
	h.values[child], h.values[parent] = h.values[parent], h.values[child]
}

func (h *Heap) parent(index int) int {
	return (index - 1) / 2
}

func IndexOf[T constraints.Ordered](val T, arr []T) int {
	for idx := range arr {
		if arr[idx] == val {
			return idx
		}
	}
	return 0
}

func isEmpty[T constraints.Ordered](val []T) bool {
	return len(val) == 0
}
