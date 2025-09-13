package minimum_path_sum

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
