# 215. Kth Largest Element

<br>

---

<br>

link: https://leetcode.com/problems/kth-largest-element-in-an-array

<br>

Given an integer array nums and an integer `k`, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

Can you solve it without sorting?

<br> 

Example 1:

```
Input: nums = [3,2,1,5,6,4], k = 2
Output: 5
```


<br>

## Coding - 1 Heap

```go

import "container/heap"

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	H := MinHeap([]interface{}{})

	for _, num := range nums {
		heap.Push(&H, num)
		if H.Len() > k {
			heap.Pop(&H)
		}
	}

	return heap.Pop(&H).(int)
}

type MinHeap []interface{}

func (h MinHeap) Len() int {
	return len(h)
}

// i priority > j
func (h MinHeap) Less(i, j int) bool {
	return h[i].(int) < h[j].(int)
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Pop() interface{} {
	ret := (*h)[h.Len()-1]
	(*h) = (*h)[:h.Len()-1]
	return ret.(int)
}

func (h *MinHeap) Push(val interface{}) {
	*h = append(*h, val)
}
```

<br>
<br>

## Coding - 2 Quick Sort

```go

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	return quickSelect(nums, 0, len(nums)-1, k)
}

func quickSelect(nums []int, start, end, k int) int {
	if start >= end {
		return nums[start]
	}

	cursor, writer := start, start

	for cursor < end {

		if nums[cursor] < nums[end] {
			// swap
			nums[cursor], nums[writer] = nums[writer], nums[cursor]
			writer++
		}

		cursor++
	}

	nums[writer], nums[end] = nums[end], nums[writer]

	targetIdx := len(nums) - k
	if writer == targetIdx {
		return nums[writer]
	} else if writer > targetIdx {
		// find left side
		return quickSelect(nums, start, writer-1, k)
	} else {
		// find right side
		return quickSelect(nums, writer+1, end, k)
	}

}
```