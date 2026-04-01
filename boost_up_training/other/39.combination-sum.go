/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	state := []int{}

	var backtracking func(idx int, currSum int)
	backtracking = func(idx int, currSum int) {
		if currSum == target {
			tmp := make([]int, len(state))
			copy(tmp, state)
			result = append(result, tmp)
			return
		} else if currSum > target {
			return
		}

		// state
		for i := idx; i < len(candidates); i++ {
			num := candidates[i]
			// 1. update state
			state = append(state, num)

			// 2. go next layer
			backtracking(i, currSum+num)

			// 3. rollback state
			state = state[:len(state)-1]
		}
	}

	backtracking(0, 0)

	return result
}

// @lc code=end

