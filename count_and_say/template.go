package count_and_say

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	dp := make([]string, n)
	dp[0] = "1"
	for i := 1; i < n; i++ {
		dp[i] = runLengthEncoding(dp[i-1])
	}
	return dp[n-1]
}

// RLE func
func runLengthEncoding(inputStr string) string {
	var result strings.Builder
	count := 1

	for i := 1; i < len(inputStr); i++ {
		if inputStr[i] == inputStr[i-1] {
			count++
		} else {
			result.WriteString(strconv.Itoa(count))
			result.WriteByte(inputStr[i-1])
			count = 1
		}
	}
	result.WriteString(strconv.Itoa(count))
	result.WriteByte(inputStr[len(inputStr)-1])
	return result.String()
}

func Go() {
	println(countAndSay(4))
}
