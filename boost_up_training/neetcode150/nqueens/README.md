# 51. N-Queens

<br>

---

<br>

## Coding

```go
func solveNQueens(n int) [][]string {
	result := [][]string{}
	board := NewBoard(n)

	var backtracking func(row int)
	backtracking = func(row int) {
		if row == n {
			// reach the end
			result = append(result, board.snapshot())
			return
		}

		for col := 0; col < n; col++ {
			if board.isSafe(row, col) {
				// update state
				board.put(row, col)
				// go next layer
				backtracking(row + 1)
				// roll back state
				board.remove(row, col)
			}
		}
	}

	backtracking(0)

	return result
}

type Board struct {
	n            int
	cols         []bool
	diaginal     map[int]bool
	aitiDiaginal map[int]bool
	state        [][]bool
}

func NewBoard(n int) *Board {
	state := make([][]bool, n)
	for i := range state {
		state[i] = make([]bool, n)
	}

	return &Board{
		n:            n,
		cols:         make([]bool, n),
		diaginal:     make(map[int]bool),
		aitiDiaginal: make(map[int]bool),
		state:        state,
	}
}

func (b *Board) isSafe(row, col int) bool {
	// check cols
	if b.cols[col] {
		return false
	}

	diaginalKey := row - col
	aitiDiaginalKey := row + col

	// check diaginal & aitiDiaginal
	if b.diaginal[diaginalKey] || b.aitiDiaginal[aitiDiaginalKey] {
		return false
	}

	return true
}

func (b *Board) put(row, col int) {
	diaginalKey := row - col
	aitiDiaginalKey := row + col

	b.cols[col] = true
	b.diaginal[diaginalKey] = true
	b.aitiDiaginal[aitiDiaginalKey] = true
	b.state[row][col] = true
}

func (b *Board) remove(row, col int) {
	diaginalKey := row - col
	aitiDiaginalKey := row + col

	b.cols[col] = false
	b.diaginal[diaginalKey] = false
	b.aitiDiaginal[aitiDiaginalKey] = false
	b.state[row][col] = false
}

func (b *Board) snapshot() []string {
	res := make([]string, b.n)

	for row := 0; row < b.n; row++ {

		str := ""
		for col := 0; col < b.n; col++ {
			// ".Q..","...Q","Q...","..Q."]
			if b.state[row][col] {
				str += "Q"
			} else {
				str += "."
			}
		}

		res[row] = str
	}

	return res
}
```

<br>
<br>

## Time & Space Complexity

```
Time: O(n square)

board.check(), board.put(), board.remove() is O(1) so I ignore it.
so the backtracking maximum run (n square) times.

although board.snapshot() also take (n square) times. but it only run several time (answer len([][]string) times).

Space: O(n) 

max backtracking call stack is n.
```

<br>
<br>

## AI Feedback

### Time Complexity: O(N!) You estimated O(N²), but it is actually much larger: O(N!) (N-factorial).

* The Backtracking Tree: In the first row, you have $N$ possible columns to place a queen. In the second row, you have at most $N-1$ safe columns. In the third row, you have at most $N-2$ safe columns, and so on.

* The Math: This means the maximum number of recursive calls you will make is roughly $N \times (N-1) \times (N-2) \dots \times 1$, which equals $N!$.

* The Snapshot: While board.snapshot() does take O(N²), you only run it when a valid solution is found. The number of valid solutions is strictly less than the number of nodes in the exploration tree.

**Conclusion**: The dominant factor is the size of the backtracking search space, leaving you with an overall worst-case time complexity of O(N!).

<br>
<br>

## Space Complexity: O(N²)

### You estimated O(N), but because of how you track the board state, it is actually O(N²).

* Call Stack: You are correct that the recursive call stack goes exactly $N$ levels deep, taking O(N) space.

* Maps and Arrays: Your cols array, diaginal map, and aitiDiaginal map store at most $N$ or $2N$ keys, which takes O(N) space.

* The 2D State Matrix (The Culprit): You initialized state [][]bool as an $N \times N$ matrix. Storing an $N \times N$ grid requires O(N²) auxiliary space.

**Conclusion**: Because the 2D matrix is the largest data structure you are holding in memory during the run, the overall space complexity becomes O(N²). (Note: Complexity analysis typically excludes the space required for the final output result array).