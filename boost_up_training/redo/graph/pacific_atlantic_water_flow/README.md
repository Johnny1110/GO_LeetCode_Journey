# 417. Pacific Atlantic Water Flow

<br>

---

<br>

link: https://leetcode.com/problems/pacific-atlantic-water-flow/

<br>

* Top-Left -> [0, ?] & [?, 0]
* Bottom-Right -> [m, ?] & [?, n]

<br>

There is an m x n rectangular island that borders both the Pacific Ocean and Atlantic Ocean. The Pacific Ocean touches the island's left and top edges, and the Atlantic Ocean touches the island's right and bottom edges.

The island is partitioned into a grid of square cells. You are given an m x n integer matrix heights where heights[r][c] represents the height above sea level of the cell at coordinate (r, c).

The island receives a lot of rain, and the rain water can flow to neighboring cells directly north, south, east, and west if the neighboring cell's height is less than or equal to the current cell's height. Water can flow from any cell adjacent to an ocean into the ocean.

Return a 2D list of grid coordinates result where result[i] = [ri, ci] denotes that rain water can flow from cell (ri, ci) to both the Pacific and Atlantic oceans.

<br>

## Coding - 1

```go
type State int

const (
	UNVISITED State = iota
	TOP_LEFT_OK
	BOT_RIGHT_OK
	BOTH_OK
)

func pacificAtlantic(heights [][]int) [][]int {
	result := [][]int{}
	m, n := len(heights), len(heights[0])

	stateTracking := make([][]State, m)
	for idx := range m {
		stateTracking[idx] = make([]State, n)
	}

	var dfsFromTopLeft func(preHight, row, col int) bool
	dfsFromTopLeft = func(preHight, row, col int) bool {

		if col >= n || col < 0 || row >= m || row < 0 {
			return false
		}

		if stateTracking[row][col] == TOP_LEFT_OK || stateTracking[row][col] == BOTH_OK {
			return true
		}

		thisHight := heights[row][col]
		if preHight <= thisHight {
			// OK
			if stateTracking[row][col] == BOT_RIGHT_OK {
				stateTracking[row][col] = BOTH_OK

				result = append(result, []int{row, col})
			} else {
				stateTracking[row][col] = TOP_LEFT_OK
			}
			// explore 4 path
			dfsFromTopLeft(thisHight, row-1, col)
			dfsFromTopLeft(thisHight, row+1, col)
			dfsFromTopLeft(thisHight, row, col-1)
			dfsFromTopLeft(thisHight, row, col+1)
		}

		return true
	}

	var dfsFromBotRight func(preHight, row, col int) bool
	dfsFromBotRight = func(preHight, row, col int) bool {

		if col >= n || col < 0 || row >= m || row < 0 {
			return false
		}

		if stateTracking[row][col] == BOT_RIGHT_OK || stateTracking[row][col] == BOTH_OK {
			return true
		}

		thisHight := heights[row][col]
		if preHight <= thisHight {
			// OK
			if stateTracking[row][col] == TOP_LEFT_OK {
				stateTracking[row][col] = BOTH_OK

				result = append(result, []int{row, col})
			} else {
				stateTracking[row][col] = BOT_RIGHT_OK
			}
			// explore 4 path
			dfsFromBotRight(thisHight, row-1, col)
			dfsFromBotRight(thisHight, row+1, col)
			dfsFromBotRight(thisHight, row, col-1)
			dfsFromBotRight(thisHight, row, col+1)
		}

		return true
	}

	// check from top-left: [0, ?] [?, 0]
	// check from bot_right: [m, ?] [?, n]

	for idx := range n {
		dfsFromTopLeft(0, 0, idx)
		dfsFromBotRight(0, m-1, idx)
	}

	for idx := range m {
		dfsFromTopLeft(0, idx, 0)
		dfsFromBotRight(0, idx, n-1)
	}

	return result
}
```
