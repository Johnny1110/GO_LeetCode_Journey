# 347. Top K Frequent Elements


<br>

---

<br>

link: https://leetcode.com/problems/top-k-frequent-elements/description/

<br>
<br>

## Thinking

Last time we did Leetcode 215 and we're try to find Kth largest element in an array.

Now we have the very similar problem. just change from "largest value" to "largest frequency".

So first of all, we need a `numFreqMap := make(map[int]int)`, `key` is number, `value` is frequency.

Collect all elements from input array to this `numFreqMap`.

<br>

Next step, create a min-heap to PUSH all numFreqMap's key into it, the comparison factor is numFreqMap's value. 
And also we keep the heap's length equals to k, so the top element of heap always the "Kth largest frequency" number.

<br>
<br>

# Coding - Heap

```go
import (
	"container/heap"
)

type IntMinHeap struct {
	data       []int
	numFreqMap map[int]int
}

func (h *IntMinHeap) Len() int {
	return len(h.data)
}

// this is min-heap
func (h *IntMinHeap) Less(i, j int) bool {
	ip := h.numFreqMap[h.data[i]]
	jp := h.numFreqMap[h.data[j]]
	return ip < jp
}

func (h *IntMinHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *IntMinHeap) Push(x interface{}) {
	h.data = append(h.data, x.(int))
}

func (h *IntMinHeap) Pop() interface{} {
	x := h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	// key: num val: frequency
	numFreqMap := make(map[int]int)
	for _, num := range nums {
		numFreqMap[num]++
	}

	// create a min int heap
	H := &IntMinHeap{
		data:       make([]int, 0),
		numFreqMap: numFreqMap,
	}

	// push all numFreqMap into
	for num := range numFreqMap {

		heap.Push(H, num)
		if H.Len() > k {
			heap.Pop(H)
		}
	}

	return H.data
}
```

<br>

Result: 

![1.png](imgs/1.png)

<br>
<br>

## Thinking Enhancement

> What's the maximum possible frequency any number can have?

If the array has N elements, the max frequency is... N!

So we can create N+1 buckets, where bucket[freq] holds all numbers with that frequency.


### Visualizing It

```
nums = [1, 1, 1, 2, 2, 3], k = 2

Step 1: Build freqMap
{1: 3, 2: 2, 3: 1}

Step 2: Build buckets (index = frequency)
bucket[0] = []
bucket[1] = [3]      ← frequency 1
bucket[2] = [2]      ← frequency 2
bucket[3] = [1]      ← frequency 3
bucket[4] = []
bucket[5] = []
bucket[6] = []

Step 3: Walk backwards, collect K elements
bucket[3] → [1]      ← got 1 element
bucket[2] → [2]      ← got 2 elements, done!

Result: [1, 2]
```

## Coding - Bucket Sort

```go

```