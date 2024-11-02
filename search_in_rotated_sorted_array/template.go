package search_in_rotated_sorted_array

import (
	"fmt"
	"go_leetcode_journey/common"
)

func search(nums []int, target int) int {
	pivotIdx := 0

	// try to split 2 sections, find pivot.
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			pivotIdx = i
		}
	}

	startIdx := 0
	endIdx := len(nums) - 1
	if nums[0] <= target && target <= nums[pivotIdx] {
		endIdx = pivotIdx
	} else {
		startIdx = pivotIdx + 1
		endIdx = len(nums) - 1
	}

	return binSearch(nums, startIdx, endIdx, target)
}

// {4, 5, 6, 7, 0, 1, 2} 3
func binSearch(nums []int, startIdx, endIdx, target int) int {

	midIdx := startIdx + (endIdx-startIdx)/2
	if startIdx > endIdx {
		return -1
	}

	fmt.Println("start:", startIdx, "end:", endIdx, "mid:", midIdx)
	if target == nums[midIdx] {
		return midIdx
	}
	if startIdx == endIdx {
		return -1
	}
	if target < nums[midIdx] {
		return binSearch(nums, startIdx, midIdx-1, target)
	}
	if target > nums[midIdx] {
		return binSearch(nums, midIdx+1, endIdx, target)
	}

	return -1
}

// Go call this func in main.go
func Go() {
	//common.Assert_answer(search([]int{4, 5, 6, 7, 0, 1, 2}, 4), 0)
	//common.Assert_answer(search([]int{4, 5, 6, 7, 0, 1, 2}, 5), 1)
	//common.Assert_answer(search([]int{4, 5, 6, 7, 0, 1, 2}, 6), 2)
	//common.Assert_answer(search([]int{4, 5, 6, 7, 0, 1, 2}, 7), 3)
	//common.Assert_answer(search([]int{4, 5, 6, 7, 0, 1, 2}, 0), 4)
	//common.Assert_answer(search([]int{4, 5, 6, 7, 0, 1, 2}, 1), 5)
	//common.Assert_answer(search([]int{4, 5, 6, 7, 0, 1, 2}, 2), 6)
	//
	//common.Assert_answer(search([]int{4, 5, 6, 7, 0, 1, 2}, 0), 4)
	//common.Assert_answer(search([]int{4, 5, 6, 7, 0, 1, 2}, 3), -1)

	//common.Assert_answer(search([]int{1}, 0), -1)
	//common.Assert_answer(search([]int{1}, 1), 0)

	common.Assert_answer(search([]int{1, 3, 5}, 0), -1)
}
