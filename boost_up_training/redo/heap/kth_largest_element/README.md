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

## Coding - 1

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