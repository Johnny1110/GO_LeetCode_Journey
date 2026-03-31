/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start
func search(nums []int, target int) int {

	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		// fmt.Printf("left: %v, mid: %v, right: %v \n", left, mid, right)

		if nums[mid] == target {
			return mid
		}

		isLeftValid := nums[left] <= nums[mid]

		if isLeftValid {
			if target >= nums[left] && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// rightValid
			if target <= nums[right] && target >= nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

// @lc code=end

