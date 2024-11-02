package search_in_rotated_sorted_array_refactor

import (
	"go_leetcode_journey/common"
)

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	// Find the pivot where the rotation happens
	pivotIdx := findPivotIdx(nums)

	// Array is not rotated, just do a normal binary search
	if pivotIdx == -1 {
		return binSearch(nums, 0, len(nums)-1, target)
	}

	// If target is pivot himself, just return pivot idx.
	if nums[pivotIdx] == target {
		return pivotIdx
	}

	if nums[0] <= target && target <= nums[pivotIdx] {
		return binSearch(nums, 0, pivotIdx, target)
	} else {
		return binSearch(nums, pivotIdx+1, len(nums)-1, target)
	}
}

func binSearch(nums []int, startIdx, endIdx, target int) int {
	for startIdx <= endIdx {
		midIdx := startIdx + (endIdx-startIdx)/2

		if nums[midIdx] == target {
			return midIdx
		} else if nums[midIdx] > target {
			endIdx = midIdx - 1
		} else {
			startIdx = midIdx + 1
		}
	}
	return -1
}

func findPivotIdx(nums []int) int {
	start, end := 0, len(nums)-1
	for start < end {
		mid := start + (end-start)/2

		if mid < end && nums[mid] > nums[mid+1] {
			return mid
		}
		if mid > start && nums[mid] < nums[mid-1] {
			return mid - 1
		}

		if nums[start] >= nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

// Go call this func in main.go
func Go() {
	common.Assert_answer(search([]int{4, 5, 6, 7, 0, 1, 2}, 4), 0)
	common.Assert_answer(search([]int{4, 5, 6, 7, 0, 1, 2}, 5), 1)
	common.Assert_answer(search([]int{4, 5, 6, 7, 0, 1, 2}, 6), 2)
	common.Assert_answer(search([]int{4, 5, 6, 7, 0, 1, 2}, 7), 3)
	common.Assert_answer(search([]int{4, 5, 6, 7, 0, 1, 2}, 0), 4)
	common.Assert_answer(search([]int{4, 5, 6, 7, 0, 1, 2}, 1), 5)
	common.Assert_answer(search([]int{4, 5, 6, 7, 0, 1, 2}, 2), 6)

	common.Assert_answer(search([]int{4, 5, 6, 7, 0, 1, 2}, 0), 4)
	common.Assert_answer(search([]int{4, 5, 6, 7, 0, 1, 2}, 3), -1)

	common.Assert_answer(search([]int{1}, 0), -1)
	common.Assert_answer(search([]int{1}, 1), 0)

	common.Assert_answer(search([]int{1, 3, 5}, 0), -1)
}
