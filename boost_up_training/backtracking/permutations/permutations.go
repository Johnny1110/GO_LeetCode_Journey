package permutations

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	used := make([]bool, len(nums))
	backtracking([]int{}, used, nums, &result)
	return result
}

func backtracking(currentState []int, used []bool, nums []int, result *[][]int) {
	// 1. store the result
	if len(currentState) == len(nums) {
		tmp := make([]int, len(currentState))
		copy(tmp, currentState)
		*result = append(*result, tmp)
		return
	}

	// 2. for loop and select a unselected element of nums
	for i := 0; i < len(nums); i++ {
		if used[i] {
			// if used, just skip
			continue
		}

		// select an unused num
		num := nums[i]
		// mark as used
		used[i] = true
		//  update currentState
		currentState = append(currentState, num)
		// go deeper
		backtracking(currentState, used, nums, result)
		// undo state and used mark
		currentState = currentState[:len(currentState)-1]
		used[i] = false
	}
}
