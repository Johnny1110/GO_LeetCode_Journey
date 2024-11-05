package find_first_and_last_position_of_element_in_sorted_array

import "fmt"

// Go Array + BinSearch
func searchRange(nums []int, target int) []int {
	leftmost := binSearchSideMost(nums, target, true)
	rightmost := binSearchSideMost(nums, target, false)
	//rightmost := -1
	return []int{leftmost, rightmost}
}

func binSearchSideMost(nums []int, target int, leftmost bool) int {
	left, right := 0, len(nums)-1
	for left <= right {

		mid := left + (right-left)>>1

		//fmt.Println("left:", left, "right:", right, "mid:", mid)

		if nums[mid] < target {
			left = mid + 1
		}
		if nums[mid] > target {
			right = mid - 1
		}

		if nums[mid] == target {
			if left == right {
				return mid
			}

			if right-left == 1 {
				if !leftmost && nums[right] == target {
					return right
				}
				return mid
			}

			if leftmost {
				right = mid
			} else {
				left = mid
			}
		}
	}

	return -1
}

// Go call this func in main.go
func Go() {
	res1 := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	fmt.Println("res1:", res1, "answer:", "[3 4]")

	res2 := searchRange([]int{2, 2}, 2)
	fmt.Println("res2:", res2, "answer:", "[0 1]")
}
