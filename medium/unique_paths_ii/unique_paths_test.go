package unique_paths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	//Input: obstacleGrid = [[0,1],[0,0]]
	//	Output: 1
	obstacleGrid := make([][]int, 2)
	obstacleGrid[0] = []int{0, 1}
	obstacleGrid[1] = []int{0, 0}

	assert.Equal(t, 1, uniquePathsWithObstacles(obstacleGrid))
}

func Test_2(t *testing.T) {
	//Input: obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
	//	Output: 2
	obstacleGrid := make([][]int, 3)
	obstacleGrid[0] = []int{0, 0, 0}
	obstacleGrid[1] = []int{0, 1, 0}
	obstacleGrid[2] = []int{0, 0, 0}
	assert.Equal(t, 2, uniquePathsWithObstacles(obstacleGrid))
}
