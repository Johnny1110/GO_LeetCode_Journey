package cwmw

func maxArea(height []int) int {
	bestAns := 0

	pointerA, pointerB := 0, len(height)-1

	for pointerA < pointerB {
		A, B := height[pointerA], height[pointerB]
		w, h := pointerB-pointerA, min(A, B)
		bestAns = max(bestAns, w*h)

		if A > B {
			pointerB--
		} else if A < B {
			pointerA++
		} else {
			pointerB--
			pointerA++
		}
	}

	return bestAns

}
