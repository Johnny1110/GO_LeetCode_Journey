package trapping_rain_water

import (
	"go_leetcode_journey/common"
)

type BarInfo struct {
	tallestLeftBarHeight  int
	tallestRightBarHeight int
	targetBarHeight       int
}

func trap(inputs []int) int {
	result := 0
	allPositionBarInfos := make([]BarInfo, 0)

	// Init first position
	allPositionBarInfos = append(allPositionBarInfos, initFirstBar(inputs))

	for i := 1; i < len(inputs); i++ {
		currentBar := BarInfo{}
		currentBar.targetBarHeight = inputs[i]
		previousBar := allPositionBarInfos[i-1]
		currentBar.tallestLeftBarHeight = max(previousBar.tallestLeftBarHeight, previousBar.targetBarHeight)
		currentBar.tallestRightBarHeight = max(previousBar.tallestRightBarHeight, currentBar.targetBarHeight)

		if currentBar.tallestRightBarHeight == currentBar.targetBarHeight {
			tallestRight := findTallestRight(inputs, i+1)
			currentBar.tallestRightBarHeight = tallestRight
		}

		allPositionBarInfos = append(allPositionBarInfos, currentBar)

		smallestSideBar := min(currentBar.tallestLeftBarHeight, currentBar.tallestRightBarHeight)
		waiter := smallestSideBar - currentBar.targetBarHeight
		if waiter > 0 {
			result += waiter
		}

	}

	return result
}

func findTallestRight(inputs []int, idx int) int {
	maxNum := 0
	for i := idx; i < len(inputs); i++ {
		maxNum = max(maxNum, inputs[i])
	}
	return maxNum
}

func initFirstBar(inputs []int) BarInfo {
	firstBar := BarInfo{}
	firstBar.targetBarHeight = inputs[0]
	firstBar.tallestLeftBarHeight = 0

	for i := 1; i < len(inputs); i++ {
		if inputs[i] > firstBar.tallestRightBarHeight {
			firstBar.tallestRightBarHeight = inputs[i]
		}
	}

	return firstBar
}

// Go call this func in main.go
func Go() {
	example1 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	common.Assert_answer(trap(example1), 6)
}
