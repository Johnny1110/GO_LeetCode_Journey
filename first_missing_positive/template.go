package first_missing_positive

import (
	"go_leetcode_journey/common"
)

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		// ⬇ if number not in where it should be (nums[nums[i]-1]) .
		for nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			// Why using for ? because we swapped nums[i] to where it should be, then we got a new number in nums[i],
			// we got to check new number again. until not match up with above judgmental.
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
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

	example6 := []int{-50, 0, 1, 3, 5, 100}
	ans6 := firstMissingPositive(example6)
	common.Assert_answer(ans6, 2)

	example7 := []int{-1, 4, 2, 1, 9, 10}
	ans7 := firstMissingPositive(example7)
	common.Assert_answer(ans7, 3)
}
