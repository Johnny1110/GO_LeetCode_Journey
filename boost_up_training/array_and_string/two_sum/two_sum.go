package two_sum

func twoSum(nums []int, target int) []int {
	// example: [2, 7, 11, 15]
	// target: 9
	// expect: [0, 1]

	diffMap := map[int]int{}

	for idx, num := range nums {
		diffMap[target-num] = idx
	}

	// find
	for idxA, num := range nums {
		if idxB, exists := diffMap[num]; exists && idxA != idxB {
			return []int{idxA, idxB}
		}
	}

	return nil
}
