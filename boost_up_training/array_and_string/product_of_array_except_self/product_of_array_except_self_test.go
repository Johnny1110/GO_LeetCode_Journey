package product_of_array_except_self

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	//Input: nums = [1,2,3,4]
	//	Output: [24,12,8,6]
	nums := []int{1, 2, 3, 4}
	res := productExceptSelf(nums)
	assert.Equal(t, []int{24, 12, 8, 6}, res)
}

func Test_2(t *testing.T) {
	//Input: nums = [-1,1,0,-3,3]
	//	Output: [0,0,9,0,0]
	nums := []int{-1, 1, 0, -3, 3}
	res := productExceptSelf(nums)
	assert.Equal(t, []int{0, 0, 9, 0, 0}, res)
}
