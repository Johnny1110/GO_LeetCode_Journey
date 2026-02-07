# 295. Find Median from Data Stream

<br>

---

<br>

link: https://leetcode.com/problems/find-median-from-data-stream/description/

<br>
<br>

## Thinking

### Topic:

* Two Pointers
* Design
* Sorting
* Heap (Priority Queue)
* Data Stream

<br>

### Approach we can try:

1. Insertion Sort — easiest to understand
2. Two Heaps — the "aha!" solution, most elegant
3. Balanced BST — More complex to implement


<br>

---

<br>

## Insertion Sort

### The Idea

Keep a sorted slice. When a new number comes:

1. Find the correct position (binary search)
2. Insert it there

**How do you find the insertion position efficiently?**

I think we can use a array to store the income data as data stream. 
and we're using binary search to find the target position index.

**After finding the position, how do you insert in Go?**

Just like regular array insert into the middle of the data array, and move every right element forward to right side.

**What's the complexity of each operation?**

Find the position index is O(Log N), insert action is O(N)

<br>
<br>

### Coding

```go
type MedianFinder struct {
	data []int
}

func Constructor() MedianFinder {
	return MedianFinder{
		data: make([]int, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.data) == 0 {
		this.data = append(this.data, num)
		return
	}

	// find the insert position idx:
	left, right := 0, len(this.data)-1 // init start end.
	positionIdx := this.binarySearch(num, left, right)

	// perform insert:
	this.data = append(this.data, 0) // extend 1 slot.
	copy(this.data[positionIdx+1:], this.data[positionIdx:])
	this.data[positionIdx] = num
}

func (this *MedianFinder) FindMedian() float64 {
	dataLen := len(this.data)
	// odd or even:
	isOdd := dataLen%2 == 1

	if isOdd {
		return float64(this.data[dataLen/2])
	} else {
		// even
		a := this.data[dataLen/2]
		b := this.data[dataLen/2-1]
		return float64(a+b) / 2
	}
}

func (this *MedianFinder) binarySearch(num, left, right int) int {
	if left >= right {
		if this.data[left] > num {
			return left
		} else {
			return left + 1
		}
	}

	middle := (right + left) / 2 // calculate middle

	if this.data[middle] == num {
		return middle + 1
	} else if this.data[middle] > num {
		return this.binarySearch(num, left, middle-1)
	} else {
		return this.binarySearch(num, middle+1, right)
	}
}
```

![1.png](imgs/1.png)


<br>
<br>

## 