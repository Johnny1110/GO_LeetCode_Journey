# 287. Find the Duplicate Number

<br>

---

<br>

link: https://leetcode.com/problems/find-the-duplicate-number

<br>
<br>

## Thinking

I have no clue about this, I just asking for some hint.

### Claude AI Hint:

Defining a function `f(i) = nums[i]`. 

Since values are in [1, n] and indices are [0, n], each value "points to" another index.

For example with `[1, 3, 4, 2, 2]`:

```
index 0 → value 1 → index 1
index 1 → value 3 → index 3
index 3 → value 2 → index 2
index 2 → value 4 → index 4
index 4 → value 2 → index 2  ← cycle!
```

<br>

### Inference

The problem is seems like a nature linked list with cycle.

Trace through the first index-0, it will lead us go to index() 

<br>

**But wait! what if value just equals to index?**

For example with value-2 sit on index-2  `[1, 4, 2, 3, 3]`:

```

... 
index 2 → value 2 → index 2
-> infinity loop on index-2
```

Then how do first timewe get in to index 2?

index? -> value 2 -> index 2

Is it possible? yes, but it will lead us to cicle result. I think it doesn't matter.

<br>

I did linked-list detection algo before. Let's adapt it in this problem.

<br>
<br>

## Coding


```go
func findDuplicate(nums []int) int {
    // TODO
}
```