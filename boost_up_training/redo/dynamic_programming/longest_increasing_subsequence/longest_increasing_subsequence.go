package longest_increasing_subsequence

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	// better tail
	tails := []int{}
	tails = append(tails, nums[0])

	for idx := 1; idx < len(nums); idx++ {
		currentNum := nums[idx]

		targetTailPosition := binarySearch(tails, currentNum)

		if targetTailPosition == len(tails) {
			tails = append(tails, currentNum)
		} else {
			tails[targetTailPosition] = currentNum
		}

	}

	return len(tails)
}

func binarySearch(tails []int, num int) int {
	startIdx, endIdx := 0, len(tails)-1

	for startIdx <= endIdx {

		midIdx := (startIdx + endIdx) / 2

		if tails[midIdx] == num {
			return midIdx
		} else if tails[midIdx] < num {
			startIdx = midIdx + 1
		} else {
			endIdx = midIdx - 1
		}

	}

	return startIdx
}
