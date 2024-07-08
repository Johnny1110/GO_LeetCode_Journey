package three_sum

import (
	"go_leetcode_journey/common"
	"sort"
)

func threeSum(nums []int) [][]int {
	// 1. sort nums asc
	sort.Ints(nums)
	//fmt.Println("Sorted array:", nums)

	// forEach nums with i and make 2 pointer left = i + 1, right = len(nums) - 1
	for i := range nums {
		left := i + 1
		right := len(nums) - 1
	}

	return nil
}

// Go call this func in main.go
func Go() {
	// Expected Output: [[-1, -1, 2],
	//					[-1, 0, 1]]
	var ans1 [2][3]int
	ans1[0][0] = -1
	ans1[0][1] = -1
	ans1[0][2] = 2
	ans1[1][0] = -1
	ans1[1][1] = 0
	ans1[1][2] = 1
	common.Assert_answer(threeSum([]int{-1, 0, 1, 2, -1, -4}), ans1)
}
