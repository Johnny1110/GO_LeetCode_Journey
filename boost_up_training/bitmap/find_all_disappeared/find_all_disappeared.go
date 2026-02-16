package find_all_disappeared

func findDisappearedNumbers(nums []int) []int {

	// 1. mark stage
	for _, v := range nums {
		idx := abs(v) - 1

		if nums[idx] > 0 {
			nums[idx] = -nums[idx] // mark as read
		}

	}

	// 2. find positive as val
	result := make([]int, 0)
	for i, v := range nums {
		if v >= 0 {
			result = append(result, i+1)
		}
	}

	return result
}

func abs(v int) int {
	if v > 0 {
		return v
	} else {
		return -v
	}
}
