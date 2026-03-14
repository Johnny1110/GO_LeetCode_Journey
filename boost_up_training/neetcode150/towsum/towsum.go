package towsum

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for idx1, num := range nums {
		if idx2, exists := numMap[target-num]; exists {
			return []int{idx2, idx1}
		} else {
			numMap[num] = idx1
		}
	}

	return nil
}
