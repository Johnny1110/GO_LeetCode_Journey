package pacific_atlantic_water_flow

func pacificAtlantic(heights [][]int) [][]int {
	// 1. create data structure to mark which cell can flow water to target Pacific or target Atlantic.
	pacific := make([][]bool, len(heights))
	for i := range pacific {
		pacific[i] = make([]bool, len(heights[i]))
	}

	atlantic := make([][]bool, len(heights))
	for i := range atlantic {
		atlantic[i] = make([]bool, len(heights[i]))
	}

	rowLen := len(heights)
	colLen := len(heights[0])

	// 2. DFS from Pacific edges (top row + left column)
	var pacificDFS func(row, col int)
	pacificDFS = func(row, col int) {
		// 2.0 check is reachable
		if pacific[row][col] == true {
			return
		}

		// 2.1: check edge
		if row == 0 || col == 0 {
			pacific[row][col] = true
		} else {
			// 2.2: check greater than left or top
			cellHeight := heights[row][col]
			if cellHeight >= heights[row][col-1] || cellHeight >= heights[row-1][col] {
				pacific[row][col] = true
			}
		}

		// 2.3: explore to down and right
		if col+1 < colLen {
			pacificDFS(row, col+1)
		}

		if row+1 < rowLen {
			pacificDFS(row+1, col)
		}
	}

	pacificDFS(0, 0)

	// 3. DFS from Atlantic edges (bottom row + right column)
	var atlanticDFS func(row, col int)
	atlanticDFS = func(row, col int) {

		// 2.0 check is reachable
		if atlantic[row][col] == true {
			return
		}

		// 2.1: check edge
		if row == rowLen-1 || col == colLen-1 {
			atlantic[row][col] = true
		} else {
			// 2.2: check greater than right or bottom
			cellHeight := heights[row][col]
			if cellHeight >= heights[row][col+1] || cellHeight >= heights[row+1][col] {
				atlantic[row][col] = true
			}
		}

		// 2.3: explore to up and left
		if col-1 > 0 {
			atlanticDFS(row, col-1)
		}

		if row-1 > 0 {
			atlanticDFS(row-1, col)
		}
	}

	atlanticDFS(rowLen-1, colLen-1)

	// 4. Find the intersection
	res := make([][]int, 0)

	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if pacific[i][j] && atlantic[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}
