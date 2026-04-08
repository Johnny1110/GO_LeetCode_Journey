# 46. Permutations

<br>

---

<br>

## Coding

```go
func permute(nums []int) [][]int {
	result := [][]int{}
	state := make([]int, 0, len(nums))
	used := make([]bool, len(nums))

	var backtracking func()
	backtracking = func() {
		if len(state) == len(nums) {
			tmp := make([]int, len(state))
			copy(tmp, state)
			result = append(result, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			// update state
			state = append(state, nums[i])
			used[i] = true
			// go deeper
			backtracking()
			// rollback
			state = state[:len(state)-1]
			used[i] = false
		}
	}

	backtracking()

	return result
}
```

<br>
<br>

## Time & Space Complexity

```
Assume: n = length of input nums

Time: O(2 n square)

Space: O(n) -> max call stack is n. and state used is O(n)
```

<br>
<br>

## AI Feedback

### Correct Complexity: $O(n \cdot n!)$

* The "Leaves": There are $n!$ (n-factorial) possible permutations. For n = 3, that's $3 \times 2 \times 1 = 6$.
* The Work per Leaf: When you reach the base case (len(state) == len(nums)), you perform copy(tmp, state), which takes $O(n)$ time.
* The Total: Since you are generating $n!$ sequences and doing $O(n)$ work to store each one, the total time is $O(n \cdot n!)$.

> Note: $2^{n^2}$ (which you wrote) would grow significantly faster than $n!$ for large $n$, but it doesn't quite describe the branching factor here.