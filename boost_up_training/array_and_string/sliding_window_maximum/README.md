# 239. Sliding Window Maximum

<br>

---

<br>

link: https://leetcode.com/problems/sliding-window-maximum/description/

<br>

## Thinking

<br>

I think that one is pretty easy. just using 2 pointer. and I gonna init a temp to store the sum value of that k size window.

Each time moving both pointerA & B forward, add `nums[pointerB]` and also reduce `nums[pointerA-1]`

Let's try this out.

<br>
<br>

## Coding-1

```go
func maxSlidingWindow(nums []int, k int) []int {
	// k is window size, and we need calculate the first sum value of window.
	result := make([]int, len(nums)-k+1)

	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}
	result[0] = windowSum

	// make 2 pointers
	pointerA := 1
	pointerB := k

	for pointerB < len(nums) {
		// current val = previous val + pointerB - previous pointerA
		windowSum = windowSum + nums[pointerB] - abs(nums[pointerA-1])
		result[pointerA] = max(windowSum, result[pointerA-1])
		pointerA++
		pointerB++
	}

	return result
}

func abs(val int) int {
	if val < 0 {
		return -val
	} else {
		return val
	}
}
```

<br>

It's turn out that I misunderstand the problem description, I thought that was a sum problem.

I should only store the max element of the k size of window.

It's not that easy...

<br>
<br>

## Topic

* Array
* Queue
* Sliding Window
* Heap (Priority Queue)
* Monotonic Queue

<br>

How does queue work for solving this problem? That's the key point I think.

After I knew those topic, I still have no idea, let take more hint:

<br>
<br>

## Hint

1. How about using a data structure such as deque (double-ended queue)?
2. The queue size need not be the same as the window’s size.
3. Remove redundant elements and the queue should store only elements that need to be considered.

<br>
<br>

