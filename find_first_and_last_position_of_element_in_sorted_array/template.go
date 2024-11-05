package find_first_and_last_position_of_element_in_sorted_array

import "fmt"

func searchRange(nums []int, target int) []int {
	return []int{1, 1}
}

// Go call this func in main.go
func Go() {
	res1 := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	fmt.Println("res1:", res1, "answer:", "[3, 4]")
}
