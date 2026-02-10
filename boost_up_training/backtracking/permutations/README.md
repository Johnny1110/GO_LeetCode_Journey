# 46. Permutations

<br>

---

<br>

link: https://leetcode.com/problems/permutations/description/

<br>

Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

<br>

Example 1:

```
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
```

Example 2:

```
Input: nums = [0,1]
Output: [[0,1],[1,0]]
```

Example 3:

```
Input: nums = [1]
Output: [[1]]
```

<br>
<br>

## Thinking

<br>

Ask myself 4 problem:

1. What represents our "current state"? (what are we building?)
2. What represents "remaining choices"? (what haven't we decided yet?)
3. When do we record a valid subset?
4. What does "undo the choice" mean concretely here?

<br>

* Answer-1: currentState is a `[]int` which is store our choice of input `nums`
* Answer-2: like [1 ,2, 3], if we choice 1, there is [2, 3] left as remaining choices
* Answer-3: I think we can record a valid subset when nothing left in remaining choices 
* Answer-4: after deep diving (after recursive call to next level), currentState = currentState[: len(currentState)-1]


<br>

### The Backtracking Pattern

```
func backtracking:

    if currentState's length is full -> record currentState into result collection.
    for x:= range := unused
        update currentState currentState = append(currentState, x)
        deep diving to next decision level
        undo, reset x -> currentState = currentState[:len(currentState)-1]
```

<br>

### How do we implement unused?

Using a `[]bool` as used, to track which num already in used.

Each time when we need a unused num, we can iterate through the nums and also check the used array.

<br>
<br>

## Coding

```go
func permute(nums []int) [][]int {
	result := make([][]int, 0)
	used := make([]bool, len(nums))
	backtracking([]int{}, used, nums, &result)
	return result
}

func backtracking(currentState []int, used []bool, nums []int, result *[][]int) {
	// 1. store the result
	if len(currentState) == len(nums) {
		tmp := make([]int, len(currentState))
		copy(tmp, currentState)
		*result = append(*result, tmp)
		return
	}

	// 2. for loop and select a unselected element of nums
	for i := 0; i < len(nums); i++ {
		if used[i] {
			// if used, just skip
			continue
		}

		// select an unused num
		num := nums[i]
		// mark as used
		used[i] = true
		//  update currentState
		currentState = append(currentState, num)
		// go deeper
		backtracking(currentState, used, nums, result)
		// undo state and used mark
		currentState = currentState[:len(currentState)-1]
		used[i] = false
	}
}
```

Result:

![1.png](imgs/1.png)
