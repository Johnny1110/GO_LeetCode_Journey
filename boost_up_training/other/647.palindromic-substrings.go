/*
 * @lc app=leetcode id=647 lang=golang
 *
 * [647] Palindromic Substrings
 */

// @lc code=start
func countSubstrings(s string) int {
	res := 0

	// init DP
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	n := len(s)
	// iterate -> "aaa"
	for length := 1; length <= n; length++ { // 只看當前長度為 length 的 sub-string
		for i := 0; i+length-1 < n; i++ { // i 從 0 開始，i+j 必須等於 length-1, i+length-1 是 j 當的位置，j 沒頂到頭就可以繼續
			j := i + length - 1

			if s[i] == s[j] && (length <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				res++
			}
		}

	}

	return res
}

// @lc code=end

