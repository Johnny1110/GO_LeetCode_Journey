/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
	result := 0
	// [8,7,2,1]

	pointerA, pointerB := 0, len(height)-1

	for pointerA < pointerB {
		a, b := height[pointerA], height[pointerB]
		h, w := min(a, b), pointerB-pointerA
		result = max(result, h*w)

		if a > b {
			pointerB--
		} else if a < b {
			pointerA++
		} else {
			pointerB--
			pointerA++
		}
	}

	return result
}

// @lc code=end

