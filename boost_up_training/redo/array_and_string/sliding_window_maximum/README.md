# 239. Sliding Window Maximum

<br>

---

<br>

link: https://leetcode.com/problems/sliding-window-maximum

<br>
<br>

## Draft

```
Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
```

define a queue: `[]int` which contains index of nums

in first round, k = 3:

```
queue:

put idx-0: [0]

put idx-1: [0, 1]

put idx-2, but nums[2] is lower than nums[1], so ignore it.: [0, 1]
```

We always keep current biggest num at the end of the queue.

<br>

In each round sliding-window is moving, we should expire from the end of queue.

<br>
<br>

## Coding-1

```go

```