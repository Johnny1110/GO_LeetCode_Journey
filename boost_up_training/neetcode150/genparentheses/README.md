# 22. Generate Parentheses

<br>

---

<br>

Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

<br>

## Coding

```go
func generateParenthesis(n int) []string {
	result := []string{}

	var backtracking func(left, right int, currentState string)
	backtracking = func(left, right int, currentState string) {
		if len(currentState) == 2*n {
			result = append(result, currentState)
			return
		}

		if left > 0 {
			backtracking(left-1, right, currentState+"(")
		}

		if right > left {
			backtracking(left, right-1, currentState+")")
		}
	}

	backtracking(n, n, "")
	return result
}
```

<br>
<br>

## Time & Space Complexity

```
Assume: n = input n

Time: O(n)

Space: O(2n) -> O(n) // max recursive call stack is 2*n (result string length)
```

<br>

## AI Revise

* Time Complexity: $O\left(\frac{4^n}{\sqrt{n}}\right)$

It is definitely not O(n). To see why, look at how many valid combinations are generated as n increases:

* n = 1: 1 result ()

* n = 2: 2 results ()(), (())

* n = 3: 5 results

* n = 4: 14 results

<br>

The number of valid combinations follows a famous mathematical sequence called the Catalan numbers. The $n$-th Catalan number is given by the formula:

$$C_n = \frac{1}{n+1} \binom{2n}{n}$$


In plain English: The time complexity is exponential because you are building a massive decision tree.