package combination_sum

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	backtracking(candidates, &result, []int{}, 0, 0, target)
	return result
}

func backtracking(candidates []int, result *[][]int, currentState []int, currentSum, idx int, target int) {
	// 1. decide how to return and record
	// currentSum >= target do return
	if currentSum > target {
		return
	} else if currentSum == target {
		tmp := make([]int, len(currentState))
		copy(tmp, currentState)
		*result = append(*result, tmp)
		return
	}

	// 2. decision tree:
	for i := idx; i < len(candidates); i++ {
		// update currentState with i
		selected := candidates[i]
		currentState = append(currentState, selected)
		currentSum += selected
		backtracking(candidates, result, currentState, currentSum, i, target)
		// why backtracking i again, because candidates can be selected multi time.
		// we can select candidates[i] again in the next deeper backtracking
		currentState = currentState[:len(currentState)-1]
		currentSum -= selected
	}
}
