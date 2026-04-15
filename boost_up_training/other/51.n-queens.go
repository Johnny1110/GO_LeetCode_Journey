/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 */

// @lc code=start
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

// @lc code=end

