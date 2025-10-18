# 70. Climbing Stairs

<br>

---

<br>

link: https://leetcode.com/problems/climbing-stairs/

<br>
<br>

## Thinking

That one is classic, don't need too much thinking, fundamental DP problem for beginner.

<br>
<br>

## Coding

```go
func climbStairs(n int) int {
	dp := []int{1, 2}
	for i := 2; i < n; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	return dp[n-1]
}
```

standard:

![1.png](imgs/1.png)

<br>

Or I can try with non-list DP:

```go
func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	a := 1
	b := 2
	for i := 2; i < n; i++ {
		c := a + b
		a = b
		b = c
	}
	return b
}
```

That works also.