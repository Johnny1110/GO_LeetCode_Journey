package combination_sum

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	state := []int{}

	var backtracking func(state []int, sum, idx int)
	backtracking = func(state []int, sum, idx int) {
		if sum == target {
			stateSnapshot := make([]int, len(state))
			copy(stateSnapshot, state)
			result = append(result, stateSnapshot)
		} else if sum > target {
			return
		}

		// sum < target
		for x := idx; x < len(candidates); x++ {

			sum += candidates[x]
			state = append(state, candidates[x])

			// go next layer
			backtracking(state, sum, x)

			// roll back
			sum -= candidates[x]
			state = state[:len(state)-1]
		}

	}

	backtracking(state, 0, 0)
	return result

}
