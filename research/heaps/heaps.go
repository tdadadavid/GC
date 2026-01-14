package heaps

import (
	"slices"

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
	if h.isFull() {
		return
	}

	// insert at the current position (size)
	h.values[h.size] = value
	h.size++

	h.bubbleUp()
}

func (h *Heap) Remove(value int) {
	// to remove we have to find the rightmost leaf
	// if we don't have, we have to pick the leftmost leaf
	values, least := h.removeLeast()
	h.moveArrForward()

	h.values = slices.Insert(values, 0, least)
	h.bubbleDown()

}

func (h *Heap) moveArrForward() {
	for idx := h.size; idx > 0; idx-- {
		h.values[idx] = h.values[idx-1]
	}
}

func (h *Heap) removeLeast() ([]int, int) {
	least := h.values[h.size]
	h.values = slices.Delete(h.values, h.size, h.size+1)
	h.size--

	return h.values, least
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

func (h *Heap) bubbleDown() {
	index := h.size - 1

	root := h.values[index]

	if h.rightChild() > root {
		for index > 0 && h.values[index] < h.rightChild() { // here we implement a MaxHeap because condition is child > parent
			h.swap(index, h.rightChild())

			// bubble up
			index = h.rightChild()
		}
	} else if h.leftChild() > root {
		for index > 0 && h.values[index] < h.leftChild() { // here we implement a MaxHeap because condition is child > parent
			h.swap(index, h.leftChild())

			// bubble up
			index = h.leftChild()
		}
	}
}

func (h *Heap) isFull() bool {
	return h.size == cap(h.values)
}

func (h *Heap) swap(child, parent int) {
	h.values[child], h.values[parent] = h.values[parent], h.values[child]
}

func (h *Heap) parent(index int) int {
	return (index - 1) / 2
}

// Root.Left = (2 * index) + 2
func (h *Heap) leftChild() int {
	index := h.size - 1
	return (2 * index) + 1
}

// Root.Right =  (2 * index) + 2
func (h *Heap) rightChild() int {
	index := h.size - 1
	return (2 * index) + 2
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
