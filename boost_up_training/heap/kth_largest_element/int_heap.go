package kth_largest_element

type IntHeap struct {
	isMin bool
	data  []int
}

func NewIntHeap(min bool) *IntHeap {
	return &IntHeap{
		isMin: min,
		data:  make([]int, 0),
	}
}

func (h IntHeap) Peek() int {
	return h.data[0]
}

func (h IntHeap) Len() int {
	return len(h.data)
}

// Less Does h[i] have higher priority than h[j]?
func (h IntHeap) Less(i, j int) bool {
	if h.isMin {
		return h.data[i] < h.data[j]
	} else {
		return h.data[i] > h.data[j]
	}
}

func (h *IntHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// Push Logic: heapifyUp
// 1. append x to the end of data
// 2. compare to the parent node, Swap() if Less() == true
// 3. repeat step 2 until pushed value's priority is lower than current parent or no parent node exists.
func (h *IntHeap) Push(x int) {
	h.data = append(h.data, x)
	currentIdx := len(h.data) - 1

	for {
		parentIdx, exist := h.parentIndex(currentIdx)
		if !exist {
			break // reach the top
		}

		if h.Less(currentIdx, parentIdx) {
			// currentIdx has higher priority than parentIdx
			h.Swap(currentIdx, parentIdx)
			currentIdx = parentIdx
		} else {
			break
		}
	}
}

// Pop Logic: heapifyDown
// 1. save h[0] as pop return value
// 2. swap h[0] and h[n-1] then shrink the slice h.data = h.data[:n-1]
// 3. now we swapped a wrong val at the top, we need swap current h[0] by heapifyDown
// 4. compare current val to the left * right child, swap with the higher priority child. Stop when no child exists OR current val is smaller than both children.
func (h *IntHeap) Pop() int {
	ret := h.data[0]
	h.Swap(0, h.Len()-1)
	h.data = h.data[:h.Len()-1]

	// heapifyDown
	currentIdx := 0
	for {
		leftIdx, leftExists := h.leftChildIndex(currentIdx)
		if !leftExists {
			// left not exists means not any child.
			break
		}

		higherPriorityChildIdx := leftIdx

		rightIdx, rightExists := h.rightChildIndex(currentIdx)
		if rightExists && h.Less(rightIdx, leftIdx) {
			// right has high priority
			higherPriorityChildIdx = rightIdx
		}

		if h.Less(higherPriorityChildIdx, currentIdx) {
			h.Swap(higherPriorityChildIdx, currentIdx)
			currentIdx = higherPriorityChildIdx
		} else {
			break
		}
	}

	return ret
}

// heap parent left right child logic
func (h *IntHeap) leftChildIndex(i int) (int, bool) {
	leftChildIdx := 2*i + 1
	if leftChildIdx >= len(h.data) {
		return -1, false
	} else {
		return leftChildIdx, true
	}
}

// heap parent right child logic
func (h *IntHeap) rightChildIndex(i int) (int, bool) {
	rightChildIdx := 2 * (i + 1)
	if rightChildIdx >= len(h.data) {
		return -1, false
	} else {
		return rightChildIdx, true
	}
}

func (h *IntHeap) parentIndex(i int) (int, bool) {
	if i <= 0 {
		return -1, false
	}
	return (i - 1) / 2, true
}
