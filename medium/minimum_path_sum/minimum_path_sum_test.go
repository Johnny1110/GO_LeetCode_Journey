package minimum_path_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	//Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
	//	Output: 7
	grid := make([][]int, 3)
	grid[0] = []int{1, 3, 1}
	grid[1] = []int{1, 5, 1}
	grid[2] = []int{4, 2, 1}

	assert.Equal(t, 7, minPathSum(grid))
}

func Test_2(t *testing.T) {
	//Input: grid = [[1,2,3],[4,5,6]]
	//	Output: 12
	grid := make([][]int, 2)
	grid[0] = []int{1, 2, 3}
	grid[1] = []int{4, 5, 6}

	assert.Equal(t, 12, minPathSum(grid))
}
