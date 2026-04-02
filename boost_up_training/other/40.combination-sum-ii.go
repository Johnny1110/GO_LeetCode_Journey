/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	state := []int{}

	slices.Sort(candidates)

	var backtracking func(idx, curr int)
	backtracking = func(idx, curr int) {

		if curr == target {
			tmp := make([]int, len(state))
			copy(tmp, state)
			result = append(result, tmp)
			return
		}

		for i := idx; i < len(candidates); i++ {
			num := candidates[i]

			if i > idx && num == candidates[i-1] {
				// skip duplicate
				continue
			}

			if num+curr > target {
				// early break (candidates is ordered)
				break
			}

			// update state
			state = append(state, num)
			// go next layer
			backtracking(i+1, curr+num)
			// rollback
			state = state[:len(state)-1]
		}
	}

	backtracking(0, 0)

	return result
}

// @lc code=end

