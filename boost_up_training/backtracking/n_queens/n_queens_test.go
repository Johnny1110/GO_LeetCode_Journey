package n_queens

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected [][]string
	}{
		{
			name: "n=1",
			n:    1,
			expected: [][]string{
				{"Q"},
			},
		},
		{
			name:     "n=2 no solution",
			n:        2,
			expected: [][]string{},
		},
		{
			name:     "n=3 no solution",
			n:        3,
			expected: [][]string{},
		},
		{
			name: "n=4",
			n:    4,
			expected: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := solveNQueens(tt.n)
			if !compareBoards(result, tt.expected) {
				t.Errorf("solveNQueens(%d) = %v; want %v", tt.n, result, tt.expected)
			}
		})
	}
}

// compareBoards compares two slices of boards (order-independent)
func compareBoards(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	// Convert each board to a single string for comparison
	toStrings := func(boards [][]string) []string {
		result := make([]string, len(boards))
		for i, board := range boards {
			result[i] = strings.Join(board, "|")
		}
		sort.Strings(result)
		return result
	}

	aStrs := toStrings(a)
	bStrs := toStrings(b)

	for i := range aStrs {
		if aStrs[i] != bStrs[i] {
			return false
		}
	}
	return true
}

func TestSolveNQueensCount(t *testing.T) {
	// Known number of solutions for N-Queens problem
	expectedCounts := map[int]int{
		1: 1,
		2: 0,
		3: 0,
		4: 2,
		5: 10,
		6: 4,
		7: 40,
		8: 92,
	}

	for n, expected := range expectedCounts {
		t.Run(string(rune('0'+n)), func(t *testing.T) {
			result := solveNQueens(n)
			if len(result) != expected {
				t.Errorf("solveNQueens(%d) returned %d solutions; want %d", n, len(result), expected)
			}
		})
	}
}

func TestSolveNQueensValidBoards(t *testing.T) {
	for n := 1; n <= 8; n++ {
		t.Run(string(rune('0'+n)), func(t *testing.T) {
			results := solveNQueens(n)
			for i, board := range results {
				if !isValidBoard(board, n) {
					t.Errorf("solveNQueens(%d) solution %d is invalid: %v", n, i, board)
				}
			}
		})
	}
}

// isValidBoard checks if a board configuration is valid
func isValidBoard(board []string, n int) bool {
	if len(board) != n {
		return false
	}

	queens := make([][2]int, 0, n)

	// Find all queen positions
	for r, row := range board {
		if len(row) != n {
			return false
		}
		queenCount := 0
		for c, ch := range row {
			if ch == 'Q' {
				queenCount++
				queens = append(queens, [2]int{r, c})
			} else if ch != '.' {
				return false
			}
		}
		if queenCount != 1 {
			return false
		}
	}

	// Check no two queens attack each other
	for i := 0; i < len(queens); i++ {
		for j := i + 1; j < len(queens); j++ {
			r1, c1 := queens[i][0], queens[i][1]
			r2, c2 := queens[j][0], queens[j][1]

			// Same row (shouldn't happen due to check above)
			if r1 == r2 {
				return false
			}
			// Same column
			if c1 == c2 {
				return false
			}
			// Same diagonal
			if abs(r1-r2) == abs(c1-c2) {
				return false
			}
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Test_SafeZoneChecker(t *testing.T) {
	checker := NewNQueensContext(4)

	mockCurrentState := [][]bool{
		{false, true, false, false},
		{false, false, false, true},
		{true, false, false, false},
		{false, false, true, false},
	}

	checker.currentState = mockCurrentState

	checker.collectCurrentState2Result()
	fmt.Println(checker.result)
}
