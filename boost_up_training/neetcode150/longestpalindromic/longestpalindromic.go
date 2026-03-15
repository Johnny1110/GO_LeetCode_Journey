package longestpalindromic

func longestPalindrome(s string) string {

	// define DP -> index: palindromic len, value
	dp := make([][]bool, len(s)) // i ~ j (from idx-i to idx-j is palindrome)
	for idx := 0; idx < len(dp); idx++ {
		dp[idx] = make([]bool, len(s))
		dp[idx][idx] = true
	}

	return ""
}
