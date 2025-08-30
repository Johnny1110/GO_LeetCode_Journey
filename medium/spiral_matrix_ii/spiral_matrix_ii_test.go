package spiral_matrix_ii

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	//Input: n = 3
	//Output: [[1,2,3],[8,9,4],[7,6,5]]
	res := generateMatrix(3)
	expected := [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}
	assert.Equal(t, expected, res)
}

func Test_2(t *testing.T) {
	//Input: n = 1
	//Output: [[1]]
	res := generateMatrix(1)
	expected := [][]int{{1}}
	assert.Equal(t, expected, res)
}

func Test_3(t *testing.T) {
	//Input: n = 1
	//Output: [[1]]
	res := generateMatrix(4)
	expected := [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}}
	assert.Equal(t, expected, res)
}

func Test_4(t *testing.T) {
	//Input: n = 1
	//Output: [[1]]
	res := generateMatrix(5)
	fmt.Println(res)
}
