package spiralmatrix

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	cnt := m * n
	result := make([]int, cnt)

	row, col := 0, 0
	for idx := 0; idx < cnt; idx++ {
		result[idx] = matrix[row][col]
	}

	return result
}

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)
