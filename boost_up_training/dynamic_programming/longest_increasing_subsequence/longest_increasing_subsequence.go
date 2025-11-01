package longest_increasing_subsequence

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// init tails
	tails := []int{}

	// using binary-search to find tails element who is bigger than num.
	for _, num := range nums {
		idx := binarySearchIdx(&tails, num) // idx is where we can replace or append
		if idx == len(tails) {
			tails = append(tails, num)
		} else {
			tails[idx] = num
		}
	}

	return len(tails)
}

// binarySearchIdx find first element's idx bigger than num.
func binarySearchIdx(tailsPtn *[]int, num int) int {
	tails := *tailsPtn
	left, right := 0, len(tails)-1

	for left <= right {
		midIdx := (left + right) / 2
		midVal := tails[midIdx]
		if midVal > num {
			right = midIdx - 1
		} else if midVal < num {
			left = midIdx + 1
		} else {
			return midIdx
		}
	}

	return left
}
