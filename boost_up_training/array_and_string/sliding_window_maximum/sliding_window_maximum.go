package sliding_window_maximum

func maxSlidingWindow(nums []int, k int) []int {
	// k is window size, and we need calculate the first sum value of window.
	result := make([]int, len(nums)-k+1)

	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}
	result[0] = windowSum

	// make 2 pointers
	pointerA := 1
	pointerB := k

	for pointerB < len(nums) {
		// current val = previous val + pointerB - previous pointerA
		windowSum = windowSum + nums[pointerB] - abs(nums[pointerA-1])
		result[pointerA] = max(windowSum, result[pointerA-1])
		pointerA++
		pointerB++
	}

	return result
}

func abs(val int) int {
	if val < 0 {
		return -val
	} else {
		return val
	}
}
