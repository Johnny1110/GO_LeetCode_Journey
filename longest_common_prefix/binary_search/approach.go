package binary_search

import (
	"math"
	"strings"
)

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// find the minLen of all strs
	minLen := 200 // desc said longest str will be 200.
	for _, str := range strs {
		minLen = int(math.Min(float64(minLen), float64(len(str))))
	}

	// now the binary search range is 1 ~ minLen
	binLow := 1
	binHigh := minLen

	for binLow <= binHigh {
		middle := (binLow + binHigh) / 2
		if isCommonPrefix(strs, middle) {
			binLow = middle + 1
		} else {
			binHigh = middle - 1
		}
	}

	return strs[0][:(binLow+binHigh)/2]
}

func isCommonPrefix(strs []string, length int) bool {
	prefix := strs[0][:length]
	for _, str := range strs[1:] {
		if !strings.HasPrefix(prefix, str[:length]) {
			return false
		}
	}
	return true
}
