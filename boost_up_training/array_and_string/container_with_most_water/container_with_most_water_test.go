package container_with_most_water

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	//Input: height = [1,8,6,2,5,4,8,3,7]
	//	Output: 49
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	assert.Equal(t, 49, maxArea(height))
}

func Test_2(t *testing.T) {
	//Input: height = [1,8,6,2,5,4,8,3,7]
	//	Output: 49
	height := []int{1, 1}
	assert.Equal(t, 1, maxArea(height))
}
