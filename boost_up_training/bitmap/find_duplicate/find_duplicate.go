package find_duclicate

func findDuplicate(nums []int) int {
	n := len(nums) - 1
	return binarySearch(nums, 1, n)
}

func binarySearch(nums []int, left, right int) int {
	if left >= right {
		return left
	}

	mid := (left + right) / 2

	lessEqMidCount := 0
	for _, v := range nums {
		if v <= mid {
			lessEqMidCount++
		}
	}

	// shrink the bounder
	if lessEqMidCount <= mid {
		// duplicate at right side
		return binarySearch(nums, mid+1, right)
	} else {
		// duplicate at left side
		return binarySearch(nums, left, mid)
	}
}
