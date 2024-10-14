package find_the_index_of_the_first_occurrence_in_a_string

import (
	"go_leetcode_journey/common"
)

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	needleIdx := 0
	pointerA := 0
	for pointerB := 0; pointerB < len(haystack); pointerB++ {
		if haystack[pointerB] == needle[needleIdx] {
			needleIdx++
			if needleIdx == len(needle) {
				return pointerA
			}
		} else {
			pointerA++
			pointerB = pointerA - 1
			needleIdx = 0
		}
	}
	return -1
}

// Go call this func in main.go
func Go() {
	common.Assert_answer(strStr("sadbutsad", "sad"), 0)
	common.Assert_answer(strStr("leetcode", "leet"), 0)
	common.Assert_answer(strStr("leetcode", "leett"), -1)
	common.Assert_answer(strStr("leetcode", "code"), 4)
	common.Assert_answer(strStr("mississippi", "issip"), 4)
}
