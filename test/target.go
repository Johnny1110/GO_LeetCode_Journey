package test

import (
	"fmt"
	"strconv"
)

// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according
//to the following rules:

// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

// Input:
// [["5","3",".",".","7",".",".",".","."]
// ,["6",".",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]
// Output: true

//Input: board =
// [["8","3",".",".","7",".",".",".","."]
// ,["6",".",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]
// Output: false

// Time Complexity: O(n)
// Space Complexity: O(n)

type State struct {
	rows [][]bool // i: which row, j: which number
	cols [][]bool
	boxs [][]bool
}

func (s *State) checkAndUpdate(row, col int, val string) bool {
	if val == "." {
		return true
	}

	intVal, _ := strconv.Atoi(val)

	// check row
	if s.rows[row][intVal] {
		return false
	}

	// check col
	if s.cols[col][intVal] {
		return false
	}

	// check box
	boxIdx := calculateBoxIdx(row, col)
	if s.boxs[boxIdx][intVal] {
		return false
	}

	// update
	s.rows[row][intVal] = true
	s.cols[col][intVal] = true
	s.boxs[boxIdx][intVal] = true

	return true
}

// ca TODO
func calculateBoxIdx(row, col int) int {

	if row <= 2 {
		if col <= 2 {
			return 0
		} else if col <= 5 {
			return 1
		} else {
			return 2
		}
	}

	if row <= 5 {
		if col <= 2 {
			return 3
		} else if col <= 5 {
			return 4
		} else {
			return 5
		}
	}

	if row <= 8 {
		if col <= 2 {
			return 6
		} else if col <= 5 {
			return 7
		} else {
			return 8
		}
	}

	return -1
}

// rows [][]bool // i: which row, j: which number
//
//	cols [][]bool
//	boxs [][]bool
func NewState() *State {
	rows := make([][]bool, 9)
	cols := make([][]bool, 9)
	boxs := make([][]bool, 9)

	for i := range 9 {
		rows[i] = make([]bool, 10)
		cols[i] = make([]bool, 10)
		boxs[i] = make([]bool, 10)
	}

	return &State{
		rows: rows,
		cols: cols,
		boxs: boxs,
	}

}

func checkSudoku(board [][]string) bool {

	state := NewState()

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {

			if !state.checkAndUpdate(i, j, board[i][j]) {
				return false
			}

		}
	}

	return true
}

func main() {
	input := [][]string{
		{"5", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}
	result := checkSudoku(input)
	fmt.Printf("result: %v \n", result)
}
