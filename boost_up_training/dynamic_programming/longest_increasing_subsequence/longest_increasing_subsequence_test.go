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

func Test_4(t *testing.T) {
	nums := []int{4, 10, 4, 3, 8, 9}
	assert.Equal(t, 3, lengthOfLIS(nums))
}

// --------------------------------------------------------

func Test_binary_search(t *testing.T) {
	tails := []int{1, 3, 5, 7, 9, 10}
	ans_1 := binarySearchIdx(&tails, 3)
	assert.Equal(t, 1, ans_1)

	ans_2 := binarySearchIdx(&tails, 6)
	assert.Equal(t, 3, ans_2)

	ans_3 := binarySearchIdx(&tails, 11)
	assert.Equal(t, 6, ans_3)
}
