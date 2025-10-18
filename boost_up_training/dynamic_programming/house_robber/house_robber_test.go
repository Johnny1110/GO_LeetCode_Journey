package house_robber

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	// Input: nums = [1,2,3,1]
	// Output: 4
	nums := []int{1, 2, 3, 1}
	assert.Equal(t, 4, rob(nums))
}

func Test_2(t *testing.T) {
	// Input: nums = [2,7,9,3,1]
	// Output: 12
	nums := []int{2, 7, 9, 3, 1}
	assert.Equal(t, 12, rob(nums))
}

func Test_3(t *testing.T) {
	// Input: nums = [2,7,9,3,1]
	// Output: 12
	nums := []int{2, 1, 1, 2}
	assert.Equal(t, 4, rob(nums))
}
