/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxV := nums[0]
	dp = nums[0]

	for i := 1; i < len(nums); i++ {
		prevSum := dp
		curr := nums[i]

		if prevSum < 0 {
			dp = curr
		} else {
			dp = prevSum + curr
		}

		maxV = max(maxV, dp)
	}

	return maxV
}

// @lc code=end

