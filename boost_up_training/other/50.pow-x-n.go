/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	sign := n > 0
	n = abs(n)

	x = pow(x, n)
	if !sign {
		return 1 / x
	}
	return x
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	left := 1.0
	for n > 1 {
		if n%2 == 1 {
			left *= x // left = 4
		}

		n = n / 2
		x *= x
	}

	return x * left
}

// @lc code=end

