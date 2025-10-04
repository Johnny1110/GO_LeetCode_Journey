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

## Claude AI Hint:

### using a deque!

**The key insight: What information do we actually need to keep?**

* Window size k = 3
* Current window contains: [1, 3, 2]

Here's the critical question: Do we need to remember the 1?

The answer is NO! Why? Because:

* The 3 is larger than 1
* The 3 came AFTER the 1
* So the 3 will stay in the window at least as long as the 1

This means the 1 can NEVER be the maximum of any future window. Once a larger element comes after a smaller one (within the window), the smaller one becomes useless.

<br>
<br>

**Now, what about the 2 after the 3?**

Should we keep [3, 2] or just [3]?

We SHOULD keep both! Why? Because when 3 slides out of the window, 2 might become the maximum. The 2 is smaller but came later, so it might be useful in the future.

<br>

**The pattern emerges: [Monotonic Decreasing Queue(單調序列)](https://medium.com/%E6%8A%80%E8%A1%93%E7%AD%86%E8%A8%98/%E6%BC%94%E7%AE%97%E6%B3%95%E7%AD%86%E8%A8%98%E7%B3%BB%E5%88%97-monotonic-stack-queue-5ad1c35a3dfe)**

What we want to maintain is a monotonically decreasing sequence of elements (actually their indices) where:

* The front has the largest element (our current maximum)
* Each element is smaller than the previous one
* But importantly, they appear in order of their position in the array


<br>
<br>

[1 3 -1 -3 5 3 6 7]

```
Array: [1, 3, -1, -3, 5, 3, 6, 7], k = 3



## 1 window
[1 3 -1] -3 5 3 6 7
init monoQueue := [1, 2] # store index of array.

## 2 window
1 [3 -1 -3] 5 3 6 7
monoQueue := [1, 2] # check left: nums[1] still in window, keep it
monoQueue := [1, 2] # check right: nums[3] > nums[2] ? -3 > -1 = X, add to right
monoQueue := [1, 2, 3] # final

## 3 window
1 3 [-1 -3 5] 3 6 7
monoQueue := [1, 2, 3] # check left: nums[1] not in window, remove it
monoQueue := [2, 3]    # check right: nums[4] > nums[3] ? 5 > -3 = O, remove right
monoQueue := [2]       # check right: nums[4] > nums[2] ? 5 > -1 = O, remove right
monoQueue := []        # empty: add index into queue
monoQueue := [4]       # final


## 4 window
1 3 -1 [-3 5 3] 6 7
monoQueue := [4] # check left: nums[4] still in window, keep it
monoQueue := [4] # check right: nums[5] > nums[4] ? 3 > 5 = X, add to right
monoQueue := [4, 5] # final
 
## 5 window
1 3 -1 -3 [5 3 6] 7
monoQueue := [4, 5] # check left: nums[4] still in window, so keep it
monoQueue := [4, 5] # check right: nums[6] > nums[5] ? 6 > 3 = O, remove right
monoQueue := [4]    # check right: nums[6] > nums[4] ? 6 > 5 = O, remove right
monoQueue := []     # empty: add index into queue
monoQueue := [6]    # final

## 6 window
1 3 -1 -3 5 [3 6 7]
monoQueue := [6]    # check left: nums[6] still in window, so keep it
monoQueue := [6]    # check right: nums[7] > nums[6] ? 7 > 6 = O, remove right
monoQueue := []     # empty: add index into queue
monoQueue := [7]    # final
```

<br>

so each window's first element is the answer index:

```
window-1: 1
window-2: 1
window-3: 4
window-4: 4
window-5: 6
window-6: 7

final result: [nums[1], nums[1], nums[4], nums[4], nums[6], nums[7]]

[3, 3, 5, 5, 6, 7]
```

<br>

Now I understand, Let's implement this.

```go

```