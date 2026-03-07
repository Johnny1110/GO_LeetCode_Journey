# 295. Find Median from Data Stream

<br>

---

<br>

link: https://leetcode.com/problems/find-median-from-data-stream

<br>
<br>

## Coding - 1 (Sorting)

```go
type MedianFinder struct {
	nums []int
}

func Constructor() MedianFinder {
	return MedianFinder{
		nums: []int{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	// binsearch position
	targetIdx := binarySearch(this.nums, num)

	if targetIdx == len(this.nums) {
		this.nums = append(this.nums, num)
	} else {
		// insert
		this.nums = append(this.nums[:targetIdx+1], this.nums[targetIdx:]...)
		this.nums[targetIdx] = num
	}

}

func (this *MedianFinder) FindMedian() float64 {
	length := len(this.nums)
	if length%2 == 0 {
		// even
		numA := this.nums[length/2]
		numB := this.nums[(length/2)-1]
		return float64(numA+numB) / 2
	} else {
		// odd
		return float64(this.nums[length/2])
	}
}

func binarySearch(nums []int, num int) int {

	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == num {
			return mid + 1
		} else if nums[mid] > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
```

<br>
<br>

## Coding - 2 (2 heap)

```go
import "container/heap"

type MedianFinder struct {
	left  IntHeap
	right IntHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		left: IntHeap{
			min:  false,
			data: []interface{}{},
		},
		right: IntHeap{
			min:  true,
			data: []interface{}{},
		},
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.left, num)

	val := heap.Pop(&this.left)
	heap.Push(&this.right, val)

	if this.right.Len() > this.left.Len()+1 {
		// balance
		val = heap.Pop(&this.right)
		heap.Push(&this.left, val)
	}

}

func (this *MedianFinder) FindMedian() float64 {
	if this.right.Len() == 0 {
		return 0
	}

	if this.left.Len() == this.right.Len() {
		leftVal, rightVal := this.left.Peek().(int), this.right.Peek().(int)
		return (float64(leftVal) + float64(rightVal)) / 2
	} else {
		return float64(this.right.Peek().(int))
	}
}

// =============================================

type IntHeap struct {
	min  bool
	data []interface{}
}

func (h IntHeap) Len() int {
	return len(h.data)
}

func (h IntHeap) Less(i, j int) bool {
	if h.min {
		return h.data[i].(int) < h.data[j].(int)
	} else {
		return h.data[i].(int) > h.data[j].(int)
	}
}

func (h IntHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *IntHeap) Push(val interface{}) {
	h.data = append(h.data, val)
}

func (h *IntHeap) Pop() interface{} {
	ret := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return ret.(int)
}

func (h IntHeap) Peek() interface{} {
	return h.data[0]
}
```

<br>
<br>

## Coding - 2 (2 heap revamp)

```go

```