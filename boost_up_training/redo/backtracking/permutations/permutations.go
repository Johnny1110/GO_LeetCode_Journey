package permutations

func permute(nums []int) [][]int {
	result := make([][]int, 0)

	state := make([]int, 0, len(nums))
	used := make(map[int]bool)

	var backtracking func(state []int, used map[int]bool)
	backtracking = func(state []int, used map[int]bool) {

		if len(state) == len(nums) {
			stateSnapshot := make([]int, len(state))
			copy(stateSnapshot, state)
			result = append(result, stateSnapshot)
			return
		}

		for _, num := range nums {
			if used[num] {
				continue
			}

			state = append(state, num)
			used[num] = true

			// go next layer
			backtracking(state, used)

			// rollback
			state = state[:len(state)-1]
			used[num] = false
		}
	}

	backtracking(state, used)

	return result
}
