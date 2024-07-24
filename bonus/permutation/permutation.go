package permutation

import "fmt"

func permutation(nums []int) [][]int {
	result := make([][]int, 0)
	backtracking(&result, &nums, 0)
	return result
}

func backtracking(result *[][]int, nums *[]int, index int) {
	if index >= len(*nums) { // if no more layer to reach
		// copy current nums permutation to answer
		numsCopy := make([]int, len(*nums))
		copy(numsCopy, *nums)
		*result = append(*result, numsCopy)
	}

	// swap
	for i := index; i < len(*nums); i++ {
		swap(*nums, index, i)
		// in first layer (index = 0): swap every number (nums[0], nums[1], nums[2]) with first nums[0].
		// in second layer (index = 1): swap every number (nums[1], nums[2]) with first nums[1].
		// in third layer (index = 2): swap every number (nums[2]) with first nums[2].
		backtracking(result, nums, index+1) // after swap, trying goes deeper until reach the bottom (index = 3).
		swap(*nums, index, i)               // swap back.
	}
}

func swap(nums []int, a int, b int) {
	// swap nums[a] & nums[b]
	nums[a], nums[b] = nums[b], nums[a]
}

func Go() {
	fmt.Println("permutation ans is: ", permutation([]int{1, 2, 3}))
}
