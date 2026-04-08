/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
	result := [][]int{}
	state := make([]int, 0, len(nums))
	used := make([]bool, len(nums))

	var backtracking func()
	backtracking = func() {
		if len(state) == len(nums) {
			tmp := make([]int, len(state))
			copy(tmp, state)
			result = append(result, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			// update state
			state = append(state, nums[i])
			used[i] = true
			// go deeper
			backtracking()
			// rollback
			state = state[:len(state)-1]
			used[i] = false
		}
	}

	backtracking()

	return result
}

// @lc code=end

