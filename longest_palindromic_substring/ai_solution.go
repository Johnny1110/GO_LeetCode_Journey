package longest_palindromic_substring

// 使用中心扩展法查找最长回文子串
func _longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	var start, end int
	for i := 0; i < len(s); i++ {
		len1 := _expandAroundCenter(s, i, i)   // 奇数长度回文串
		len2 := _expandAroundCenter(s, i, i+1) // 偶数长度回文串
		maxLen := max(len1, len2)
		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

// 辅助函数：以left和right为中心向两边扩展，返回回文串的长度
func _expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

// 辅助函数：返回两个数中较大的数
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
