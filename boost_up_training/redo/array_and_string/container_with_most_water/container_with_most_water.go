package container_with_most_water

func maxArea(height []int) int {
    
	pointerA, pointerB := 0, len(height)-1
	tmp := 0

	for pointerA < pointerB {
		w := pointerB-pointerA
		h := min(height[pointerA], height[pointerB]) // using smaller one
		tmp = max(tmp, w*h)

		// moving smaller one
		if height[pointerA] > height[pointerB] {
			pointerB--
		} else {
			pointerA++
		}
	}

	return tmp
}