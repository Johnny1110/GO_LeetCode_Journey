package subsets

func subsets(nums []int) [][]int {
	// build backtracking structure.
	// what do we need? resultCollection, currentState, remainChoices, undoAction
	resultCollection := make([][]int, 0)
	backtrackingA(&resultCollection, []int{}, nums, 0)
	return resultCollection
}

func backtrackingA(result *[][]int, currentState []int, nums []int, idx int) {
	// 1. record currentState as result
	tmp := make([]int, len(currentState))
	copy(tmp, currentState)
	*result = append(*result, tmp)

	// 2. for loop
	for i := idx; i < len(nums); i++ {
		// 3. update currentState
		currentState = append(currentState, nums[i])
		// 4. deep diving
		backtrackingA(result, currentState, nums, i+1)
		// 5. undo state
		currentState = currentState[:len(currentState)-1]
	}
}

// nums + idx = remainChoices
func backtrackingB(resultCollection *[][]int, currentState []int, nums []int, idx int) {
	if idx == len(nums) {
		// reach the end, store currentState as result.
		tmp := make([]int, len(currentState))
		copy(tmp, currentState)
		*resultCollection = append(*resultCollection, tmp)
		return
	}

	currentState = append(currentState, nums[idx])             // update currentState
	backtrackingB(resultCollection, currentState, nums, idx+1) // deep dive into next decision level
	currentState = currentState[:len(currentState)-1]          // undo currentState
	backtrackingB(resultCollection, currentState, nums, idx+1) // depp dive with other decision
}
