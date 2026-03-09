# Subsets

<br>

---

<br>

link: https://leetcode.com/problems/subsets

<br>

Given an integer array nums of unique elements, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.

<br>


Example 1:

```
Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
```

Example 2:

```
Input: nums = [0]
Output: [[],[0]]
```

<br>
<br>

## Coding - 1

```go
func subsets(nums []int) [][]int {

	result := [][]int{}

	var backtracking func(idx int, state []int)
	backtracking = func(idx int, state []int) {
		// 1. save result
		stateSnapshot := make([]int, len(state))
		copy(stateSnapshot, state)
		result = append(result, stateSnapshot)

		for i := idx; i < len(nums); i++ {
			// 2. change state
			state = append(state, nums[i])
			// go deeper:
			backtracking(i+1, state)

			// recover state
			state = state[:len(state)-1]
		}
	}

	backtracking(0, []int{})

	return result
}
```