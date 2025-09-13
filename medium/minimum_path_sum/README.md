# 64. Minimum Path Sum

<br>

---

<br>

link: https://leetcode.com/problems/minimum-path-sum/description/

<br>

## Thinking

<br>

I guess it's kinda like previous problem unique path. so it means this is another DP problem.

so I may using 2D array to solve this problem. now I should define the DP array.

`dp[i][j]` represent the minimum path sum (the result).

let try it out.

<br>
<br>

## Coding

```go
func minPathSum(grid [][]int) int {

	for rowIdx, row := range grid {
		for columnIdx := range row {

			topSum := -1
			leftSum := -1

			if rowIdx > 0 {
				topSum = grid[rowIdx-1][columnIdx]
			}
			if columnIdx > 0 {
				leftSum = grid[rowIdx][columnIdx-1]
			}

			grid[rowIdx][columnIdx] = grid[rowIdx][columnIdx] + minSum(topSum, leftSum)
		}
	}

	m := len(grid)
	n := len(grid[0])
	return grid[m-1][n-1]
}

func minSum(sum1 int, sum2 int) int {
	if sum1 == -1 && sum2 == -1 {
		return 0
	}
	if sum1 == -1 {
		return sum2
	} else if sum2 == -1 {
		return sum1
	} else {
		return min(sum1, sum2)
	}
}
```