package longest_increasing_subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	//Input: nums = [10,9,2,5,3,7,101,18]
	//Output: 4
	//Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.

	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	assert.Equal(t, 4, lengthOfLIS(nums))
}

func Test_2(t *testing.T) {
	nums := []int{0, 1, 0, 3, 2, 3}
	assert.Equal(t, 4, lengthOfLIS(nums))
}

func Test_3(t *testing.T) {
	nums := []int{7, 7, 7, 7, 7}
	assert.Equal(t, 1, lengthOfLIS(nums))
}
