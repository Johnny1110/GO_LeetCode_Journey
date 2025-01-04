package first_missing_positive

import (
	"go_leetcode_journey/common"
	"sort"
)

func firstMissingPositive(nums []int) int {
	nums = append(nums, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		expectNext := nums[i] + 1
		if expectNext < nums[i+1] && expectNext > 0 {
			return expectNext
		}
	}
	return nums[len(nums)-1] + 1
}

// Go call this func in main.go
func Go() {
	example1 := []int{1, 2, 3, 4, 5}
	ans1 := firstMissingPositive(example1)
	common.Assert_answer(ans1, 6)

	example2 := []int{1, 2, 4, 5}
	ans2 := firstMissingPositive(example2)
	common.Assert_answer(ans2, 3)

	example3 := []int{1, 3, 0, -1, 5, 0}
	ans3 := firstMissingPositive(example3)
	common.Assert_answer(ans3, 2)

	example4 := []int{3, 4, 1, -1}
	ans4 := firstMissingPositive(example4)
	common.Assert_answer(ans4, 2)

	example5 := []int{7, 8, 9, 11, 12}
	ans5 := firstMissingPositive(example5)
	common.Assert_answer(ans5, 1)
}
