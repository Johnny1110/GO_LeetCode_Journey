/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	farest := 0
	currentEnd := 0
	jumps := 0

	for i := 0; i < len(nums); i++ {
		steps := nums[i]
		farest = max(farest, i+steps)

		if i == currentEnd {
			jumps++
			currentEnd = farest
		}

		if currentEnd >= len(nums)-1 {
			break
		}
	}

	return jumps
}

// @lc code=end

