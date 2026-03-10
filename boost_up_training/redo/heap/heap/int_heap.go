package heap

type Heap[T any] struct {
	data []T
	less func(a, b T) bool
}

func (h Heap[T]) Len() int           { return len(h.data) }
func (h Heap[T]) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h Heap[T]) Less(i, j int) bool { return h.less(h.data[i], h.data[j]) }

// Push heapifyUp
func (h *Heap[T]) Push(val T) {
	h.data = append(h.data, val)
	currentIdx := h.Len() - 1

	for {
		if parentIdx, exists := h.getParentIdx(currentIdx); exists {
			if h.Less(parentIdx, currentIdx) {
				h.Swap(parentIdx, currentIdx)
				currentIdx = parentIdx
			} else {
				break
			}
		} else {
			break
		}
	}
}

// Pop heapyifyDown
func (h *Heap[T]) Pop() T {
	ret := h.data[0]
	// swap head and toe
	h.Swap(0, h.Len()-1)

	// shrink
	h.data = h.data[:h.Len()-1]

	currentIdx := 0
	// heapifyDown
	for {
		leftIdx, exists := h.getLeftChildIdx(currentIdx)
		if !exists {
			break
		}

		higherPriorityChildIdx := leftIdx

		if rightIdx, exists := h.getRightChildIdx(currentIdx); exists && h.Less(higherPriorityChildIdx, rightIdx) {
			higherPriorityChildIdx = rightIdx
		}

		if h.Less(currentIdx, higherPriorityChildIdx) {
			h.Swap(currentIdx, higherPriorityChildIdx)
			currentIdx = higherPriorityChildIdx
		} else {
			break
		}
	}

	return ret
}

func NewHeap[T any](less func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		data: []T{},
		less: less,
	}
}

func (h Heap[T]) getParentIdx(nodeIdx int) (int, bool) {
	if nodeIdx == 0 {
		return -1, false
	}

	return (nodeIdx - 1) / 2, true

}

func (h Heap[T]) getLeftChildIdx(nodeIdx int) (int, bool) {
	idx := nodeIdx*2 + 1
	if idx < len(h.data) {
		return idx, true
	}
	return -1, false
}

func (h Heap[T]) getRightChildIdx(nodeIdx int) (int, bool) {
	idx := nodeIdx*2 + 2
	if idx < len(h.data) {
		return idx, true
	}
	return -1, false
}
