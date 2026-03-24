# 17. Letter Combinations of a Phone Number

<br>

---

<br>

## Coding

```go
var digAlphaMap = map[uint8][]uint8{
	'2': []uint8{'a', 'b', 'c'},
	'3': []uint8{'d', 'e', 'f'},
	'4': []uint8{'g', 'h', 'i'},
	'5': []uint8{'j', 'k', 'l'},
	'6': []uint8{'m', 'n', 'o'},
	'7': []uint8{'p', 'q', 'r', 's'},
	'8': []uint8{'t', 'u', 'v'},
	'9': []uint8{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	result := []string{}

	if digits == "" {
		return result
	}

	state := make([]uint8, len(digits))

	var backtracking func(idx int)
	backtracking = func(idx int) {
		if idx >= len(digits) {
			result = append(result, string(state))
			return
		}

		alphaSet := digAlphaMap[digits[idx]]

		for _, char := range alphaSet {
			// update state
			state[idx] = char
			// go next layer
			backtracking(idx + 1)
			// rollback state
			// no need
		}
	}

	backtracking(0)

	return result
}
```

<br>
<br>

## Time & Space Complexity

```
Assume: n = length of input digits

Time: O(n 3rd power) each char of digits have at least 3 branch.

Space: O(n 3rd power) each recursive call need a space.
```

<br>
<br>

## AI Feedback

Here is the breakdown of the actual complexity:

### Time Complexity: $O(4^n \cdot n)$

* The Branching Factor: Each digit maps to either 3 or 4 letters. In the worst case (using digits '7' or '9'), every digit has 4 branches. This creates a state space of $4^n$ combinations.

* The String Construction: For each of those $4^n$ combinations, you perform a string(state) conversion, which takes $O(n)$ time to copy the bytes into a new string for the result slice.

* Total: $O(n \cdot 4^n)$.

<br>

### Space Complexity: $O(n)$

* Recursion Stack: The depth of your recursion is equal to the number of digits $n$. Therefore, the stack space is $O(n)$.

* State Array: Your state slice also takes $O(n)$ space.

* Note on Output: Usually, we don't count the result slice as part of the "auxiliary space" complexity. If you did count the output, it would be $O(n \cdot 4^n)$, but strictly speaking, the memory required to run the algorithm is $O(n)$.