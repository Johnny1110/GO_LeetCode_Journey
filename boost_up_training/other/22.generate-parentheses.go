/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func generateParenthesis(n int) []string {
	result := []string{}

	var backtracking func(left, right int, currentState string)
	backtracking = func(left, right int, currentState string) {
		if len(currentState) == 2*n {
			result = append(result, currentState)
			return
		}

		if left > 0 {
			backtracking(left-1, right, currentState+"(")
		}

		if right > left {
			backtracking(left, right-1, currentState+")")
		}
	}

	backtracking(n, n, "")
	return result
}

// @lc code=end

