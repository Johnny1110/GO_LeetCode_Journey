/*
 * @lc app=leetcode id=52 lang=golang
 *
 * [52] N-Queens II
 */

// @lc code=start
func totalNQueens(n int) int {
	state := NewChessBoard(n)
	cnt := 0

	var backtracking func(row int)
	backtracking = func(row int) {
		if row == n {
			cnt++
			return
		}

		for col := 0; col < n; col++ {

			if !state.IsSafe(row, col) {
				continue
			}

			state.Put(row, col)
			// go deeper
			backtracking(row + 1)
			// retreat
			state.Remove(row, col)
		}
	}

	backtracking(0)

	return cnt
}

// ============================

type ChessBoard struct {
	vertical     []bool
	diagonal     map[int]bool
	antiDiagonal map[int]bool
}

func NewChessBoard(n int) *ChessBoard {
	return &ChessBoard{
		vertical:     make([]bool, n),
		diagonal:     make(map[int]bool),
		antiDiagonal: make(map[int]bool),
	}
}

func (b ChessBoard) IsSafe(row, col int) bool {
	if b.vertical[col] {
		return false
	}
	if b.diagonal[row-col] {
		return false
	}
	if b.antiDiagonal[row+col] {
		return false
	}
	return true
}

func (b *ChessBoard) Put(row, col int) {
	b.vertical[col] = true
	b.diagonal[row-col] = true
	b.antiDiagonal[row+col] = true
}

func (b *ChessBoard) Remove(row, col int) {
	b.vertical[col] = false
	b.diagonal[row-col] = false
	b.antiDiagonal[row+col] = false
}

// @lc code=end

