/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	result := [][]int{}

	if len(nums) < 2 {
		return result
	}

	// sort
	slices.Sort(nums)

	for i := 0; i < len(nums)-2; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue // skip duplicate
		}

		if nums[i] > 0 { // we need a negative val
			return result
		}

		targetSum := -nums[i]

		// using 2 pointer
		pointerA, pointerB := i+1, len(nums)-1
		for pointerA < pointerB {

			A, B := nums[pointerA], nums[pointerB]
			tmpSum := A + B

			if tmpSum == targetSum {
				// found 1 result
				result = append(result, []int{nums[i], A, B})
				pointerA++
				for pointerA < pointerB && nums[pointerA] == nums[pointerA-1] {
					pointerA++
				}
				pointerB--
				for pointerA < pointerB && nums[pointerB] == nums[pointerB+1] {
					pointerB--
				}
			} else if tmpSum < targetSum {
				// we need a bigger num, move pointerA
				pointerA++
				for pointerA < pointerB && nums[pointerA] == nums[pointerA-1] {
					pointerA++
				}
			} else {
				// we need a smaller num, move pointerB
				pointerB--
				for pointerA < pointerB && nums[pointerB] == nums[pointerB+1] {
					pointerB--
				}
			}

		}
	}

	return result
}

// @lc code=end

