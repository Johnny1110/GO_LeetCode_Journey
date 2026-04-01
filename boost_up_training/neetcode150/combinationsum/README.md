# 39. Combination Sum

<br>

---

<br>

## Coding 

```go
func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	state := []int{}

	var backtracking func(idx int, currSum int)
	backtracking = func(idx int, currSum int) {
		if currSum == target {
			tmp := make([]int, len(state))
			copy(tmp, state)
			result = append(result, tmp)
			return
		} else if currSum > target {
			return
		}

		// state
		for i := idx; i < len(candidates); i++ {
			num := candidates[i]
			// 1. update state
			state = append(state, num)

			// 2. go next layer
			backtracking(i, currSum+num)

			// 3. rollback state
			state = state[:len(state)-1]
		}
	}

	backtracking(0, 0)

	return result
}
```

<br>

## Time & Space Complexity

```
Assume: n = input candidates length, h = max result []int length

Time: 0(n)

Space: O(h)
```

<br>

## AI Feedback

### Time Complexity: $O(n^h)$ ❌ Incorrect (It's not $O(n)$)

Backtracking is essentially a Depth-First Search (DFS) on an execution tree. It explores all possible valid paths, which results in exponential, rather than linear, time complexity.

Here is the breakdown for the time complexity:

* Branching Factor: At each level of the recursion tree, you can make at most $n$ choices (since your loop goes from idx to the end of the candidates array).

* Tree Depth: As established, the maximum depth of this recursion tree is $h$.

* Tree Size: In the absolute worst-case scenario (a full n-ary tree), the number of nodes in the execution tree is bounded by $O(n^h)$.

* Copying Cost: Every time you hit currSum == target (a leaf node), you execute copy(tmp, state). Copying an array of length $h$ takes $O(h)$ time.

Therefore, a loose upper bound for the time complexity is $O(n^h)$ (or more specifically, $O(h \cdot n^h)$ if we heavily factor in the array copying at the leaves).