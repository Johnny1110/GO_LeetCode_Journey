package trapping_rain_water

import "go_leetcode_journey/common"

func trap(height []int) int {
	ans := towPointersSolution(height)
	return ans
}

func towPointersSolution(height []int) int {
	pointerA := 0
	pointerB := 1
	result := 0

	for pointerB < len(height) {
		if height[pointerB] >= height[pointerA] {
		}
	}

	return result
}

// Go call this func in main.go
func Go() {
	example1 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	common.Assert_answer(trap(example1), 6)
}
