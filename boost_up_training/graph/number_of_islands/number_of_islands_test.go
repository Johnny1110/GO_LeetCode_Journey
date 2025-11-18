package number_of_islands

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	//Input: grid = [
	//	["1","1","1","1","0"],
	//	["1","1","0","1","0"],
	//	["1","1","0","0","0"],
	//	["0","0","0","0","0"]
	//]
	//Output: 1

	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	result := numIslands(grid)
	fmt.Println("result:", result)
	assert.Equal(t, 1, result)
}

func Test_2(t *testing.T) {
	//Input: grid = [
	//	["1","1","0","0","0"],
	//	["1","1","0","0","0"],
	//	["0","0","1","0","0"],
	//	["0","0","0","1","1"]
	//]
	//Output: 3

	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	result := numIslands(grid)
	fmt.Println("result:", result)
	assert.Equal(t, 3, result)
}
