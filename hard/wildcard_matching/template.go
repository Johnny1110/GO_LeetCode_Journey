package wildcard_matching

import "go_leetcode_journey/common"

func isMatch(s string, p string) bool {
	// TODO
	return false
}

// Go call this func in main.go
func Go() {
	common.Assert_answer(isMatch("aa", "a"), false)
	common.Assert_answer(isMatch("a", "a"), true)
	common.Assert_answer(isMatch("aa", "a?"), true)
	common.Assert_answer(isMatch("aa", "a*"), true)
	common.Assert_answer(isMatch("aaaaab", "a*b"), true)
	common.Assert_answer(isMatch("aaaaabbbbb", "a*b"), false)
	common.Assert_answer(isMatch("aaaaabbbbb", "a*b*"), true)
	common.Assert_answer(isMatch("aabb", "a?b*"), true)
	common.Assert_answer(isMatch("aabb", "a?b"), false)
}
