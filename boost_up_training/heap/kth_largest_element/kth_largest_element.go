package kth_largest_element

type IntMinHeap []int

func (h IntMinHeap) Len() int {
	return len(h)
}

func (h IntMinHeap) Less(i, j int) bool {
	// min-heap: smaller value = higher priority
	return h[i] < h[j]
}

func (h IntMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntMinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]     // grab last element
	*h = old[0 : n-1] // shrink slice
	return x
}

func findKthLargestByHeap(nums []int, k int) int {
	H := NewIntHeap(true)

	for _, num := range nums {
		H.Push(num)
		if H.Len() > k {
			H.Pop()
		}
	}

	// Kth Largest element will sit on the top.
	return H.Peek()
}

func quickSelectKthLargest(nums []int, start, end, k int) int {
	if start == end {
		return nums[start]
	}

	// init 2 pointers:
	i, j := start, start

	for j < end {

		if nums[j] < nums[end] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}

		j++
	}

	pivotIdx := i
	nums[pivotIdx], nums[end] = nums[end], nums[pivotIdx]

	if pivotIdx == len(nums)-k {
		return nums[pivotIdx]
	} else if pivotIdx < len(nums)-k {
		return quickSelectKthLargest(nums, pivotIdx+1, end, k)
	} else {
		return quickSelectKthLargest(nums, start, pivotIdx-1, k)
	}
}

func findKthLargestByQuickSelect(nums []int, k int) int {
	start, end := 0, len(nums)-1
	return quickSelectKthLargest(nums, start, end, k)
}

func findKthLargest(nums []int, k int) int {
	return findKthLargestByQuickSelect(nums, k)
}
