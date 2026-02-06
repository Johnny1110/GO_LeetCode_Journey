package kth_largest_element

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "basic case",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        2,
			expected: 5,
		},
		{
			name:     "k equals 1",
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:        1,
			expected: 6,
		},
		{
			name:     "k equals array length",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        6,
			expected: 1,
		},
		{
			name:     "single element",
			nums:     []int{1},
			k:        1,
			expected: 1,
		},
		{
			name:     "duplicates",
			nums:     []int{7, 10, 4, 3, 20, 15},
			k:        3,
			expected: 10,
		},
		{
			name:     "all same elements",
			nums:     []int{1, 1, 1, 1},
			k:        2,
			expected: 1,
		},
		{
			name:     "negative numbers",
			nums:     []int{-1, 2, 0},
			k:        1,
			expected: 2,
		},
		{
			name:     "large k",
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:        4,
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findKthLargest(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("findKthLargest(%v, %d) = %d; want %d", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}

func Test_Heap(t *testing.T) {
	h := &IntMinHeap{3, 1, 4, 1, 5}
	heap.Init(h) // heapify!

	heap.Push(h, 9)

	fmt.Println("min:", (*h)[0]) // peek at root

	for h.Len() > 0 {
		fmt.Println(heap.Pop(h)) // pops in ascending order
	}
}
