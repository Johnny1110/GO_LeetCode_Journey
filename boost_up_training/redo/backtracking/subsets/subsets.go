package subsets

func subsets(nums []int) [][]int {

	result := [][]int{}

	var backtracking func(idx int, state []int)
	backtracking = func(idx int, state []int) {
		// 1. save result
		stateSnapshot := make([]int, len(state))
		copy(stateSnapshot, state)
		result = append(result, stateSnapshot)

		for i := idx; i < len(nums); i++ {
			// 2. change state
			state = append(state, nums[i])
			// go deeper:
			backtracking(i+1, state)

			// recover state
			state = state[:len(state)-1]
		}
	}

	backtracking(0, []int{})

	return result
}
