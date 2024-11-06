package search_insert_position

import (
	"go_leetcode_journey/common"
)

func searchInsert(nums []int, target int) int {

	leftIndex, rightIndex := 0, len(nums)-1

	for leftIndex <= rightIndex {

		midIdx := leftIndex + (rightIndex-leftIndex)/2
		if nums[midIdx] == target {
			return midIdx
		}
		if nums[midIdx] < target {
			leftIndex = midIdx + 1
		}
		if nums[midIdx] > target {
			rightIndex = midIdx - 1
		}
	}
	if rightIndex < 0 {
		return 0
	}
	if nums[rightIndex] > target {
		return rightIndex + 2
	} else {
		return rightIndex + 1
	}
}

// Go call this func in main.go
func Go() {
	common.Assert_answer(searchInsert([]int{1, 3, 5, 6}, 5), 2)
	common.Assert_answer(searchInsert([]int{1, 3, 5, 6}, 2), 1)
	common.Assert_answer(searchInsert([]int{1, 3, 5, 6}, 7), 4)
	common.Assert_answer(searchInsert([]int{1}, 7), 1)
	common.Assert_answer(searchInsert([]int{1}, 0), 0)
}
