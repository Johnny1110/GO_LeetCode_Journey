package climbing_stairs

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	a := 1
	b := 2
	for i := 2; i < n; i++ {
		c := a + b
		a = b
		b = c
	}
	return b
}
