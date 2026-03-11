package wordsearch

func exist(board [][]byte, word string) bool {
	state := NewBoardState(board)

	var backtracking func(row, col int) bool
	backtracking = func(row, col int) bool {

		if state.wordIdx == len(word) {
			return true
		}

		if row >= len(board) || row < 0 || col >= len(board[0]) || col < 0 {
			return false
		}

		if state.used[row][col] {
			return false
		}

		thisChar := board[row][col]
		if thisChar != word[state.wordIdx] {
			return false
		}

		// update state
		state.used[row][col] = true
		state.wordIdx++

		// explore up down left right
		thisResult := backtracking(row+1, col) || backtracking(row-1, col) || backtracking(row, col+1) || backtracking(row, col-1)
		if thisResult {
			return true
		} else {
			// rollback
			state.used[row][col] = false
			state.wordIdx--
			return false
		}
	}

	for row := range len(board) {
		for col := range len(board[row]) {
			if backtracking(row, col) {
				return true
			}
		}
	}

	return false
}

// --------------------------------------------

type BoardState struct {
	used    [][]bool
	wordIdx int
}

func NewBoardState(board [][]byte) *BoardState {

	used := make([][]bool, len(board))

	for i := range len(used) {
		used[i] = make([]bool, len(board[i]))
	}

	return &BoardState{
		used:    used,
		wordIdx: 0,
	}
}
