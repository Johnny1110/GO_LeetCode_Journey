package contains_duplicate

func containsDuplicate(nums []int) bool {
	checkMap := map[int]bool{}

	for _, num := range nums {
		if _, exists := checkMap[num]; exists {
			return true
		} else {
			checkMap[num] = true
		}
	}

	return false
}
