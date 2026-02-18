# 41. First Missing Positive

<br>

---

<br>

link: https://leetcode.com/problems/first-missing-positive

<br>

Given an unsorted integer array nums. Return the smallest positive integer that is not present in nums.

You must implement an algorithm that **runs in O(n) time and uses O(1) auxiliary space.**


<br>
<br>

## Thinking

Take a example: `[7, 8, 9, 11, 12]`

```
7 : 0000 0111
8 : 0000 1000
9 : 0000 1001
11: 0000 1011
12: 0000 1100
```

We don't have to care about negative numbers.

What if we can use a `map[int]bool` to store all positive value in to it.

then we start from 1 to MAX_INT to find which is not exist in this map.

map size could be a const, because w know the len of `nums`, execute time is O(n) also. because we have 2 loops,

one is loop through nums, another is loop from 1 to the answer num.

<br>

### Claude AI Feedback

Your thinking is right — the map approach works logically. But the problem says `O(1)` auxiliary space, and a `map[int]bool` uses `O(n)` space.

Here's the key insight: **use the input array itself as the map.**


<br>

### Inference

Let me think about how to do it with input array itself.


`[7, 8, 9, 11, 12]`

<br>

If we want to using that "map" approach, what we need is a "thing" that can help us find a number is in the input array or not quickly.


What do we got?

* array got indies (0, 1, 2, 3, 4)

I don't think that works.

Maybe we can iterate through the input nums, if we find a nums[i] is **0 or negative or it is out of bound of nums range**, we just discard it (`nums[i] = -99`), otherwise, we move the number to it `nums[val]`

After that we loop through this nums again. if nums[i] != i+1, we return `i+1` as result.

<br>
<br>

## Coding


```go

func firstMissingPositive(nums []int) int {
	numsLen := len(nums)

	// stage-1. sort the nums
	for i := 0; i < numsLen; i++ {
		v := nums[i]

		if v <= 0 || v > numsLen {
			// discard: mark as -1
			nums[i] = -1
		} else {
			// do swap v to where it should be (j).
			j := v - 1

			if i == j {
				continue // already sitting in right place
			}

			nums[i], nums[j] = nums[j], nums[i]
			if nums[i] == nums[j] { // duplicate case
				nums[i] = -1 // discard nums[i]
			}
			i-- // sort nums[i] again, until it be discarded.
		}

	}

	// stage-2, find min index which val is not negative
	for i, v := range nums {
		if v != i+1 {
			return i + 1
		}
	}

	return numsLen + 1
}
```

Result:

![1](imgs/1.png)