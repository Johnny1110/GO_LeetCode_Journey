package four_sum

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	// 1. sort nums asc
	sort.Ints(nums)
	return kSum(nums, target, 0, 4)
}

func kSum(nums []int, target int, start int, k int) [][]int {
	res := make([][]int, 0)

	if start == len(nums) {
		return res
	}

	// 2. calculate avg with target and k
	avg := target / k
	// the num that we're judging, if max num is lower than avg or min num is bigger than avg, return res.
	if nums[start] > avg || nums[len(nums)-1] < avg {
		return res
	}

	if k == 2 {
		return twoSum(nums, target, start)
	}

	for i := start; i < len(nums); i++ {
		if i == start || nums[i] != nums[i-1] {
			for _, subset := range kSum(nums, target-nums[i], i+1, k-1) {
				res = append(res, append([]int{nums[i]}, subset...))
			}
		}
	}

	return res
}

func twoSum(nums []int, target int, start int) [][]int {
	res := make([][]int, 0)
	low := start
	high := len(nums) - 1

	for high > low {
		tempSum := nums[low] + nums[high]
		if tempSum < target || (low > start && nums[low] == nums[low-1]) {
			low++
		} else if tempSum > target || (high < len(nums)-1 && nums[high] == nums[high+1]) {
			high--
		} else {
			res = append(res, []int{nums[low], nums[high]})
			low++
			high--
		}
	}
	return res
}

func Go() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println("result: ", fourSum(nums, target))
	fmt.Println("answer: ", "[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]")

	fmt.Println("------------------------------------------------")

	nums2 := []int{2, 2, 2, 2, 2}
	target2 := 8
	fmt.Println("result: ", fourSum(nums2, target2))
	fmt.Println("answer: ", "[[2,2,2,2]]")
}
