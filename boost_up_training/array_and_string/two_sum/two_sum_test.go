package two_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	assert.Equal(t, []int{0, 1}, twoSum(nums, target))
}

func Test_2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	assert.Equal(t, []int{1, 2}, twoSum(nums, target))
}

func Test_3(t *testing.T) {
	nums := []int{3, 3}
	target := 6
	assert.Equal(t, []int{0, 1}, twoSum(nums, target))
}
