# 448. Find All Numbers Disappeared in an Array

<br>

---

<br>

link: https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array



<br>
<br>

## Thinking

I have no idea how to using bitmap to solve this problem.

The only approach I can deliver is using a set[int] (in golang is map[int]), put every elements into it.

After that, iterate by range `1 ~ n` find which one is not in sets. 

> Follow up: Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

The set apporach is not allow to use.

<br>

### Decompose the problem

* The range of nums index is `0 ~ n-1` 
* The range of nums value is `1 ~ n`

natural mapping: `value v → index v-1`


<br>

I think maybe we can sort the nums, `v := nums[i]` and `v` should put into `nums[v-1]`

After resort the `nums`, iterate through the `nums` again, to find which `nums[i] != i+1`.

<br>

### Claude Code Feedback

There are two valid in-place approaches here. You've described one:

* Approach 1: **Cyclic Sort** — swap each value to its "correct" position, then scan for mismatches. This works well.

* Approach 2: **Negation marking** — this is the hint I was nudging toward. Instead of moving values around, you visit index `v-1` and __negate__ the value there as a "mark." After one pass, any index still holding a positive value was never visited — meaning that index's corresponding number is missing.

<br>

I think both of em are the same, but I prefer to Approach-2.

<br>
<br>

## Coding

```go
func findDisappearedNumbers(nums []int) []int {

	// 1. mark stage
	for _, v := range nums {
		idx := abs(v) - 1

		if nums[idx] > 0 {
			nums[idx] = -nums[idx] // mark as read
		}

	}

	// 2. find positive as val
	result := make([]int, 0)
	for i, v := range nums {
		if v >= 0 {
			result = append(result, i+1)
		}
	}

	return result
}

func abs(v int) int {
	if v > 0 {
		return v
	} else {
		return -v
	}
}
```

<br>

Result:

![1](imgs/1.png)