# 39. Combination Sum

<br>

---

<br>

link: https://leetcode.com/problems/combination-sum

<br>

Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations for the given input.

 <br>

Example 1:

```
Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
```

Explanation:

```
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.
```

Example 2:

```
Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]
```

Example 3:

```
Input: candidates = [2], target = 1
Output: []
```

Constraints:

```
1 <= candidates.length <= 30
2 <= candidates[i] <= 40
All elements of candidates are distinct.
1 <= target <= 40
```

<br>

## Coding - 1

```go
func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	state := []int{}

	var backtracking func(state []int, sum, idx int)
	backtracking = func(state []int, sum, idx int) {
		if sum == target {
			stateSnapshot := make([]int, len(state))
			copy(stateSnapshot, state)
			result = append(result, stateSnapshot)
		} else if sum > target {
			return
		}

		// sum < target
		for x := idx; x < len(candidates); x++ {

			sum += candidates[x]
			state = append(state, candidates[x])

			// go next layer
			backtracking(state, sum, x)

			// roll back
			sum -= candidates[x]
			state = state[:len(state)-1]
		}

	}

	backtracking(state, 0, 0)
	return result

}
```