package sliding_window_maximum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	//Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
	//	Output: [3,3,5,5,6,7]
	res := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	assert.Equal(t, []int{3, 3, 5, 5, 6, 7}, res)
}

func Test_2(t *testing.T) {
	res := maxSlidingWindow([]int{1}, 1)
	assert.Equal(t, []int{1}, res)
}
