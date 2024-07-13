package three_sum

import (
	"sort"
)

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func threeSumClosest(nums []int, target int) int {
	// 1. sort nums
	sort.Ints(nums)
	numsLen := len(nums)
	// 2. closet result
	resultSum := nums[0] + nums[1] + nums[2]

	// 3. iterate sums
	for i := range nums {
		leftIdx := i + 1
		rightIdx := numsLen - 1

		for rightIdx > leftIdx {
			currentSum := nums[i] + nums[leftIdx] + nums[rightIdx]
			currentDiff := absInt(target - currentSum)
			lastDiff := absInt(target - resultSum)
			if lastDiff > currentDiff {
				resultSum = currentSum
			}
			if currentDiff == 0 {
				// best sol
				return resultSum
			}
			if target-currentSum > 0 {
				// increment left idx
				leftIdx++
			} else {
				// decrement right idx
				rightIdx--
			}
		}
	}

	return resultSum
}

// Go call this func in main.go
func Go() {
	//[1,1,1,0] -> target: -100 result: 2
}
