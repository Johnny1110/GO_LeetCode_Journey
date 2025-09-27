package container_with_most_water

func maxArea(height []int) int {
	result := 0

	pointerA := 0
	pointerB := len(height) - 1

	for pointerA < pointerB {
		heightA := height[pointerA]
		heightB := height[pointerB]

		temp := min(heightA, heightB) * (pointerB - pointerA)
		result = max(result, temp)

		if heightA == heightB {
			pointerA++
			pointerB--
		} else if heightA > heightB {
			// we need a bigger B, try moving B
			pointerB--
		} else {
			// we need a bigger A, try moving A
			pointerA++
		}
	}

	return result
}
