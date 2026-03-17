package reverseint

import "math"

func reverse(x int) int {

	// 31422
	sign := x > 0
	x = abs(x)
	y := 0

	for x > 0 {
		y = y*10 + x%10
		x /= 10

		if sign && y > math.MaxInt32 {
			return 0
		}
		if !sign && y > math.MaxInt32+1 {
			return 0
		}
	}

	if sign {
		return y
	} else {
		return -y
	}
}

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}
