package three_sum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	// 1. sort nums asc
	sort.Ints(nums)
	numsLen := len(nums)

	// 2. create 2D array for contain result
	ans := make([][]int, 0)

	// 3. forEach nums with i and make 2 pointer left = i + 1, right = len(nums) - 1
	for i := range nums {
		// 4. skip first elements duplicates:
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := numsLen - 1

		for left < right {
			sumResult := nums[i] + nums[left] + nums[right]

			if sumResult == 0 {
				// append nums[i], nums[left], nums[right] to ans
				ans = append(ans, []int{nums[i], nums[left], nums[right]})

				// skip left and right duplicates
				for left < right && nums[left] == nums[left+1] {
					left += 1
				}
				for left < right && nums[right] == nums[right+-1] {
					right -= 1
				}
				left++
				right--
			} else if sumResult < 0 {
				// move left idx to the right by 1
				left += 1
			} else {
				// move right idx to the left by 1
				right -= 1
			}
		}
	}

	return ans
}

// Go call this func in main.go
func Go() {
	// Expected Output: [[-1, -1, 2],
	//					[-1, 0, 1]]
	var ans1 [2][3]int
	ans1[0][0] = -1
	ans1[0][1] = -1
	ans1[0][2] = 2
	ans1[1][0] = -1
	ans1[1][1] = 0
	ans1[1][2] = 1
	//common.Assert_answer(threeSum([]int{-1, 0, 1, 2, -1, -4}), ans1)
}
