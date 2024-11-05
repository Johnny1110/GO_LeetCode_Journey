package find_first_and_last_position_of_element_in_sorted_array

import "fmt"

// Go Array + BinSearch
func searchRange(nums []int, target int) []int {
	leftmost := binSearchSideMost(nums, target, true)
	rightmost := binSearchSideMost(nums, target, false)
	return []int{leftmost, rightmost}
}

func binSearchSideMost(nums []int, target int, leftmost bool) int {
	left, right := 0, len(nums)-1
	result := -1

	for left <= right {

		mid := left + (right-left)>>1

		if nums[mid] == target {
			result = mid

			if leftmost {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	return result
}

// ---------------------------------------------------------

// searchRange Refactor
// Adjust leftmost and rightmost Searches: When searching for the leftmost or rightmost position, you don’t need to check if left == right. The binary search logic will automatically converge to the desired boundary.
// Refine the Conditionals: You don’t need the special cases where right - left == 1 since binary search with adjusted bounds will narrow in on the target.
// Shift Bitwise Operation: mid := left + (right - left) >> 1 should have parentheses around the whole calculation (i.e., (right-left)>>1). However, it’s more readable to use (right + left) / 2 in Go, which is safe for integer calculations and avoids overflow.

// ---------------------------------------------------------

// Go call this func in main.go
func Go() {
	res1 := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	fmt.Println("res1:", res1, "answer:", "[3 4]")

	res2 := searchRange([]int{2, 2}, 2)
	fmt.Println("res2:", res2, "answer:", "[0 1]")
}
