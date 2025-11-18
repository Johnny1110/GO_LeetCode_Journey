# 62. Unique Paths

<br>

---

<br>

link: https://leetcode.com/problems/unique-paths/description/

<br>
<br>


## Thinking

The character can only do 2 operation to reach the end:

* 1. Going Right
* 2. Going Down

<br>

### define DP

first of all, we assume to using a 2D array.

* `dp[i][j]` how many combination to reach `i, j`

for example:

* `dp[0][1]` = 1 : (you can only choose go RIGHT as 1 solution)
* `dp[0][2]` = 1 : (you can only choose go RIGHT 2 times as 1 solution)
* `dp[1][0]` = 1 : (you can only choose go DOWN as 1 solution)
* `dp[1][1]` = 2 : (you can choose DOWN+RIGHT or RIGHT+DOWN as 2 solutions)
...

<br>

## DP transfer formula:


* `dp[i][j] = (how many solution to reach above position) + (how many solution to reach left position)`

<br>

for example ( m = 3, n = 2):

```go
0 1 
1 2
1 3
```

result = 3

<br>

* formula: `dp[i][j] = dp[i-1][j] + dp[i][j-1]`

<br>

## Optimization

Actually, we can only using a 1D array to achieve the solution.

* define a 1D `dp[]`
* formula: `dp[i] = dp[i-1] + dp[i]`
* result is `dp[n]`

<br>
<br>

## Coding

```go
// uniquePaths
func uniquePaths(m int, n int) int {
	// 1. define 1D array as DP
	dp := make([]int, n)
	// 2. init dp (skip first, from idx: 1 to idx n-1)
	for i := 0; i < n; i++ {
		// first row always be 1 (only go right side, 1 option, no other choice)
		dp[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j-1] + dp[j]
		}
	}

	return dp[n-1]
}
```

<br>

Result: ![1.png](imgs/1.png)

<br>


