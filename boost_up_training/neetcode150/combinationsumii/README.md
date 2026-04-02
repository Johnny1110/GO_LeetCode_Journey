# 40. Combination Sum II

<br>

---

<br>

Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target.

Each number in candidates may only be used once in the combination.

Note: The solution set must not contain duplicate combinations.

<br>

Example 1:
```
Input: candidates = [10,1,2,7,6,1,5], target = 8
Output: 
    [
        [1,1,6],
        [1,2,5],
        [1,7],
        [2,6]
    ]
```

<br>
<br>

## Coding

```go
func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	state := []int{}

	slices.Sort(candidates)
	visited := make(map[string]bool)

	var backtracking func(idx, curr int)
	backtracking = func(idx, curr int) {

		if curr == target {
			tmp := make([]int, len(state))
			copy(tmp, state)

			key := intarr2str(tmp)
			if !visited[key] {
				result = append(result, tmp)
				visited[key] = true
			}

			return
		}

		for i := idx; i < len(candidates); i++ {
			num := candidates[i]
			thisSum := curr + num
			if thisSum > target {
				continue
			}

			// update state
			state = append(state, num)
			// go next layer
			backtracking(i+1, thisSum)
			// rollback
			state = state[:len(state)-1]
		}
	}

	backtracking(0, 0)

	return result
}

func intarr2str(content []int) string {
	str := ""
	for _, num := range content {
		str += strconv.Itoa(num)
	}
	return str
}
```

<br>
<br>

## Coding - Revamp

```go
func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	state := []int{}

	slices.Sort(candidates)

	var backtracking func(idx, curr int)
	backtracking = func(idx, curr int) {

		if curr == target {
			tmp := make([]int, len(state))
			copy(tmp, state)
			result = append(result, tmp)
			return
		}

		for i := idx; i < len(candidates); i++ {
			num := candidates[i]

			if i > idx && num == candidates[i-1] {
				// skip duplicate
                continue
			}

			if num+curr > target {
                // early break (candidates is ordered)
				break
			}

			// update state
			state = append(state, num)
			// go next layer
			backtracking(i+1, curr+num)
			// rollback
			state = state[:len(state)-1]
		}
	}

	backtracking(0, 0)

	return result
}
```

<br>
<br>

## Time & Space Complexity

```
Assume: n input candidates length, h = biggest element length in answer array. 

Time: O(log n) + O(n h square) -> slice.Sort() is O(log n), iterate through candidates and recursive call is O(n h square)

Space: O(h) -> backtracking max recursive call stack is h
```

<br>

## AI Feedback

### Time Complexity: $O(2^n \times n)$ (Worst Case)

Your estimation of $O(\log n) + O(n h^2)$ is unfortunately incorrect for a couple of reasons:

1. Sorting: Go's slices.Sort() uses a highly optimized algorithm (pdqsort) which takes $O(n \log n)$ time, not $O(\log n)$.

2. Recursion Tree: Backtracking algorithms generally do not run in polynomial time like $O(n h^2)$. In the absolute worst-case scenario, you are fundamentally exploring a decision tree for $n$ elements (essentially "pick" or "don't pick" for each element). Even with early stopping and deduplication, the theoretical upper bound of paths explored is proportional to $O(2^n)$.

3. Copying to Result: Every time you find a valid combination, you run copy(tmp, state). If the combination length is $h$ (where the worst-case is $h = n$), this deep copy operation takes $O(n)$ time.

Therefore, the overall time complexity is bounded by the sorting plus the backtracking tree: $O(n \log n + 2^n \times n)$, which simplifies to the dominant term of $O(2^n \times n)$.
