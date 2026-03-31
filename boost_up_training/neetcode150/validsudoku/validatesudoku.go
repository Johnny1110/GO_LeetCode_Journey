package validsudoku

import (
	"fmt"
	"strconv"
)

func isValidSudoku(board [][]byte) bool {
	state := NewState()

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {

			if board[row][col] == '.' {
				continue
			}

			num, _ := strconv.Atoi(string(board[row][col]))
			ok := state.ValidateAndUpdate(row, col, num)
			fmt.Printf("row:%v, col:%v, validate: %v\n", row, col, ok)
			if !ok {
				return false
			}

		}
	}

	return true
}

type State struct {
	rows    [][]bool
	cols    [][]bool
	subboxs [][]bool
}

func NewState() *State {
	rows := make([][]bool, 9)
	cols := make([][]bool, 9)
	subboxs := make([][]bool, 9)

	for idx := range rows {
		rows[idx] = make([]bool, 9)
	}
	for idx := range cols {
		cols[idx] = make([]bool, 9)
	}
	for idx := range subboxs {
		subboxs[idx] = make([]bool, 9)
	}

	return &State{
		rows:    rows,
		cols:    cols,
		subboxs: subboxs,
	}
}

func (s *State) ValidateAndUpdate(row, col, num int) bool {
	numIdx := num - 1
	// 1. check row and col
	if s.rows[row][numIdx] || s.cols[col][numIdx] {
		return false
	}

	// 2. check subbox
	boxIdx := calculateSubboxIdx(row, col)
	fmt.Printf("boxIdx: %v \n", boxIdx)
	if s.subboxs[boxIdx][numIdx] {
		return false
	}

	// 3. update occupy state
	s.rows[row][numIdx], s.cols[col][numIdx], s.subboxs[boxIdx][numIdx] = true, true, true
	return true
}

func calculateSubboxIdx(row, col int) int {
	factor := row / 3
	base := factor * 3
	cv := col / 3
	return base + cv
}
