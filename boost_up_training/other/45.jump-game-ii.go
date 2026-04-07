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

	jumps := 0
	currentEnd := 0
	farthest := 0

	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])

		if i == currentEnd {
			jumps++
			currentEnd = farthest
			if currentEnd >= len(nums)-1 {
				break
			}
		}
	}

	return jumps
}

// @lc code=end

