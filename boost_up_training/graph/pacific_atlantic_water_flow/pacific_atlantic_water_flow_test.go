package pacific_atlantic_water_flow

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	heights := [][]int{
		{1},
	}

	res := pacificAtlantic(heights)

	expect := [][]int{
		{0, 0},
	}

	assert.ElementsMatch(t, res, expect)
}

func Test_2(t *testing.T) {
	heights := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}

	res := pacificAtlantic(heights)

	expect := [][]int{
		{0, 4},
		{1, 3},
		{1, 4},
		{2, 2},
		{3, 0},
		{3, 1},
		{4, 0},
	}

	assert.ElementsMatch(t, res, expect)
}
