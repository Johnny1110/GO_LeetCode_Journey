package four_sum

import "fmt"

func fourSum(nums []int, target int) [][]int {
	// TODO
	return nil
}

func Go() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println("result: ", fourSum(nums, target))
	fmt.Println("answer: ", "[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]")

	fmt.Println("------------------------------------------------")

	nums2 := []int{2, 2, 2, 2, 2}
	target2 := 8
	fmt.Println("result: ", fourSum(nums2, target2))
	fmt.Println("answer: ", "[[2,2,2,2]]")
}
