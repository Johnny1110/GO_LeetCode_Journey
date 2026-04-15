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
	diaginal     []bool
	aitiDiaginal []bool
	state        []int
}

func NewBoard(n int) *Board {
	return &Board{
		n:            n,
		cols:         make([]bool, n),
		diaginal:     make([]bool, 2*n),
		aitiDiaginal: make([]bool, 2*n),
		state:        make([]int, n),
	}
}

func (b *Board) isSafe(row, col int) bool {
	// check cols
	if b.cols[col] {
		return false
	}

	diaginalKey := row - col + b.n
	aitiDiaginalKey := row + col

	// check diaginal & aitiDiaginal
	if b.diaginal[diaginalKey] || b.aitiDiaginal[aitiDiaginalKey] {
		return false
	}

	return true
}

func (b *Board) put(row, col int) {
	diaginalKey := row - col + b.n
	aitiDiaginalKey := row + col

	b.cols[col] = true
	b.diaginal[diaginalKey] = true
	b.aitiDiaginal[aitiDiaginalKey] = true
	b.state[row] = col
}

func (b *Board) remove(row, col int) {
	diaginalKey := row - col + b.n
	aitiDiaginalKey := row + col

	b.cols[col] = false
	b.diaginal[diaginalKey] = false
	b.aitiDiaginal[aitiDiaginalKey] = false
}

func (b *Board) snapshot() []string {
	res := make([]string, b.n)
	chars := make([]uint8, b.n)

	for row := 0; row < b.n; row++ {

		queenAtCol := b.state[row]

		for col := 0; col < b.n; col++ {
			if col == queenAtCol {
				chars[col] = 'Q'
			} else {
				chars[col] = '.'
			}
		}

		res[row] = string(chars)
	}

	return res
}

// @lc code=end

