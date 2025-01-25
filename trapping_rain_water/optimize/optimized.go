package trapping_rain_water_optimize

import "go_leetcode_journey/common"

func trap(inputs []int) int {

	length := len(inputs)
	leftMax := make([]int, length)
	rightMax := make([]int, length)

	leftMax[0] = inputs[0]
	for i := 1; i < length; i++ {
		leftMax[i] = max(leftMax[i-1], inputs[i])
	}

	rightMax[length-1] = inputs[length-1]
	for i := length - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], inputs[i])
	}

	sum := 0
	for i := 0; i < length; i++ {
		sum += min(leftMax[i], rightMax[i]) - inputs[i]
	}

	return sum
}

// Go call this func in main.go
func Go() {
	example1 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	common.Assert_answer(trap(example1), 6)
}
