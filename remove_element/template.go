package remove_element

import (
	"go_leetcode_journey/common"
)

func removeElement(nums []int, val int) int {
	//return approachA(nums, val)
	return approachB(nums, val)
}

// approachA two pointer each from nums's begin(left) and the end(right).
func approachA(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	pointerA := 0
	pointerB := len(nums) - 1

	for pointerA <= pointerB {
		for nums[pointerB] == val {
			pointerB--
			if pointerB == -1 {
				break
			}
		}

		if pointerA > pointerB {
			break
		}

		if nums[pointerA] == val {
			nums[pointerA] = nums[pointerB]
			pointerB--
		}
		pointerA++
	}

	return pointerA
}

// approachB two pointer both from nums's begin (left).
func approachB(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	pointerA := 0
	for pointerB := 0; pointerB < len(nums); pointerB++ {
		if nums[pointerB] != val {
			nums[pointerA] = nums[pointerB]
			pointerA++
		}
	}
	return pointerA
}

func Go() {
	demo1 := []int{1, 2, 3}
	ans1 := removeElement(demo1, 1)
	common.Assert_answer(ans1, 2)

	demo2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	ans2 := removeElement(demo2, 2)
	common.Assert_answer(ans2, 5)

	demo3 := []int{1}
	ans3 := removeElement(demo3, 1)
	common.Assert_answer(ans3, 0)

	demo4 := []int{1, 1, 1, 1}
	ans4 := removeElement(demo4, 1)
	common.Assert_answer(ans4, 0)

	demo5 := []int{3, 2, 2, 3}
	ans5 := removeElement(demo5, 3)
	common.Assert_answer(ans5, 2)

	demo6 := []int{1, 2, 3, 4}
	ans6 := removeElement(demo6, 3)
	common.Assert_answer(ans6, 3)
}
