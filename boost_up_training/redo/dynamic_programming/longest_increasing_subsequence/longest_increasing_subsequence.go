package longest_increasing_subsequence

// Input: nums = [10,9,2,5,3,7,101,18]
// Output: 4
// Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// define do: better tails
	tails := []int{nums[0]}

	for _, num := range nums {
		positionIdx := binarySearch(tails, num)
		if positionIdx == len(tails) {
			tails = append(tails, num)
		} else {
			tails[positionIdx] = num
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
