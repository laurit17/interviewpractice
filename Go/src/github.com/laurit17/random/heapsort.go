package heapsort

import ()

type DataSequence interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type HeapSorter struct {
	list           *DataSequence
	sortedBoundary int
}

// heapindex transforms an index in a list to
// a heap index and vice versa.
func (h *HeapSorter) heapIndex(i int) int {
	return list.Len() - 1 - i
}

// gives you the list indices of the children
// of a list index.
func (h *HeapSorter) children(i int) {
	child1 := h.heapIndex(i) * 2
	child2 := h.heapindex(i)*2 + 1

	return h.heapIndex(child1), h.heapIndex(child2)
}

func (h *HeapSorter) parent(i int) {
	heapLoc := h.heapIndex(i)
	return h.heapIndex(heapLoc / 2)
}

func (h *HeapSorter) hasParent(i int) bool {
	return h.parent(i) != i
}

func (h *HeapSorter) inTheHeap(i int) bool {
	return i >= h.sortedBoundary && i < h.list.Len()
}

func (h *HeapSorter) isLeaf(i int) bool {
	leftChild, rightChild := h.children(i)
	return !h.inTheHeap(leftChild) && !h.inTheHeap(rightChild)
}

func (h *HeapSorter) min(i, j int) int {
	if list.Less(i, j) {
		return i
	}
	return j
}

func (h *HeapSorter) bubbleDown(root int) {
	if h.isLeaf(root) {
		return
	}

	leftChild, rightChild := h.children(root)
	var minChild int
	if !h.inTheHeap(leftChild) {
		minChild = rightChild
	} else if !h.inTheHeap(rightChild) {
		minChild = leftChild
	} else {
		minChild = h.min(leftChild, rightChild)
	}

	if h.min(minChild, root) == minChild {
		h.list.Swap(minChild, root)
		h.bubbleDown(minChild)
	}
}

func (h *HeapSorter) heapify(root int) {
	if h.isLeaf(root) || !h.inTheHeap(root) {
		return
	}

	leftChild, rightChild := h.children(root)
	h.heapify(leftChild)
	h.heapify(rightChild)
	var minChild int
	if !h.inTheHeap(leftChild) {
		minChild = rightChild
	} else if !h.inTheHeap(rightChild) {
		minChild = leftChild
	} else {
		minChild = min(leftChild, rightChild)
	}

	if min(minChild, root) == minChild {
		h.list.Swap(minChild, root)
	}

}

func (h *HeapSorter) Sort() *DataSequence {
	h.sortedBoundary = 0
	h.heapify(h.list.Len() - 1)
	for h.sortedBoundary = 0; h.sortedBoundary < h.list.Len(); h.sortedBoundary++ {
		h.list.Swap(h.sortedBoundary, h.list.Len()-1)
		h.bubbleDown(h.list.Len() - 1)
	}

	return h.list
}
