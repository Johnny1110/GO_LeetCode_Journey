package kth_largest_element

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	return quickSelect(nums, 0, len(nums)-1, k)
}

func quickSelect(nums []int, start, end, k int) int {
	if start >= end {
		return nums[start]
	}

	cursor, writer := start, start

	for cursor < end {

		if nums[cursor] < nums[end] {
			// swap
			nums[cursor], nums[writer] = nums[writer], nums[cursor]
			writer++
		}

		cursor++
	}

	nums[writer], nums[end] = nums[end], nums[writer]

	targetIdx := len(nums) - k
	if writer == targetIdx {
		return nums[writer]
	} else if writer > targetIdx {
		// find left side
		return quickSelect(nums, start, writer-1, k)
	} else {
		// find right side
		return quickSelect(nums, writer+1, end, k)
	}

}
