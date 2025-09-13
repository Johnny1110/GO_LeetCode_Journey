package unique_paths

// uniquePaths
func uniquePaths(m int, n int) int {
	totalSteps := m - 1 + n - 1
	return c(totalSteps, m-1)
}

func c(n int, k int) int {
	// C(n,k) = n! / (k! × (n-k)!)
	// 利用對稱性：C(n,k) = C(n,n-k)
	if k > n-k {
		k = n - k
	}

	result := 1
	for i := 0; i < k; i++ {
		result = result * (n - i) / (i + 1)
	}
	return result
}
