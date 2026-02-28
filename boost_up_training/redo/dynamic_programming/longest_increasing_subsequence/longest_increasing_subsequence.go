package longest_increasing_subsequence

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	tails := []int{nums[0]}

	for idx := 1; idx < len(nums); idx++ {

		num := nums[idx]
		tailPosition := binarySearch(tails, num)

		if tailPosition == len(tails) {
			tails = append(tails, num)
		} else {
			tails[tailPosition] = num
		}

	}

	return len(tails)
}

func binarySearch(tails []int, num int) int {
	left, right := 0, len(tails)-1

	for left <= right {
		mid := (left + right) / 2

		if tails[mid] == num {
			return mid
		} else if tails[mid] > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
