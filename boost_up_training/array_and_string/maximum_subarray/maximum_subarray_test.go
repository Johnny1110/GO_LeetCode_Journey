package maximum_subarray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := maxSubArray(nums)
	assert.Equal(t, 6, res)
}

func Test_2(t *testing.T) {
	nums := []int{1}
	res := maxSubArray(nums)
	assert.Equal(t, 1, res)
}

func Test_3(t *testing.T) {
	nums := []int{5, 4, -1, 7, 8}
	res := maxSubArray(nums)
	assert.Equal(t, 23, res)
}
