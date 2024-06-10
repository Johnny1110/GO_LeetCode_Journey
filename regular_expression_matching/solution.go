package regular_expression_matching

import "go_leetcode_journey/common"

func isMatch(s string, p string) bool {
	//return myIdea(s, p)
	//return recursiveSolution(s, p)
	return dynamicProgrammingSolution(s, p)
}

func myIdea(s, p string) bool {
	if p == "*" {
		return true
	}
	// source string idx + pattern index
	sourceInx, patternInx := 0, 0

	for {
		if sourceInx == len(s) && patternInx == len(p) {
			// when both reach to the end, return true
			return true
		}

		if sourceInx == len(s) || patternInx == len(p) {
			// one of then reach to then end first, return false
			return false
		}

		// if s = p or p = '.' skip this
		if s[sourceInx] == p[patternInx] || p[patternInx] == '.' {
			sourceInx++
			patternInx++
			continue
		} else if p[patternInx] == '*' {
			// 看下一個 PATTERN 自元是否符合 source

			if sourceInx+1 == len(s) { // end of source
				if patternInx+1 == len(p) { // if pattern also reached to the end
					return true
				} else {
					return false
				}
			}

			if s[sourceInx+1] == p[patternInx+1] || p[patternInx] == '.' {
				// both next char are match then both ++
				sourceInx++
				patternInx++
			} else {
				// if not, source ++ pattern stay cool
				sourceInx++
			}
			continue
		} else {
			return false
		}

	}
}

func recursiveSolution(input, pattern string) bool {
	if len(pattern) == 0 {
		return len(input) == 0
	}
	firstMatch := len(input) != 0 && (input[0] == pattern[0] || pattern[0] == '.')
	//firstMatch := input[0] == pattern[0] || pattern[0] == '.'

	if len(pattern) >= 2 && pattern[1] == '*' {
		// try to ignore first 2 char.
		ignoreStarResult := recursiveSolution(input, pattern[2:])
		cutFirstCharResult := firstMatch && recursiveSolution(input[1:], pattern)
		return ignoreStarResult || cutFirstCharResult
	} else {
		// no star, then both cut 1 char.
		return firstMatch && recursiveSolution(input[1:], pattern[1:])
	}
}

func dynamicProgrammingSolution(input, pattern string) bool {
	// init dp array
	inputLen := len(input)
	patternLen := len(pattern)

	// dp[x][y] means (input 0~x == pattern 0~y)
	dp := make([][]bool, inputLen+1)
	for i := range dp {
		dp[i] = make([]bool, patternLen+1)
	}

	// empty input(0~0) equals to empty pattern(0~0)
	dp[0][0] = true

	// init first row: empty input input(0~0) vs pattern(0~patternLen-1)
	// what kind of pattern can match a empty string except empty pattern: the answer is like: a*, a*b*, a*b*c*
	// so, let's set them up.
	for i := 2; i <= patternLen; i++ {
		if pattern[i-1] == '*' { // start form second char (empty(0) -> first(1) -> sec(!2!))
			dp[0][i] = dp[0][i-2] // if matched: empty input vs pattern(0~i-1) == empty input vs pattern(0~i-2)
		}
	}

	// dynamic func fill the dp:
	for dpInputIdx := 1; dpInputIdx <= inputLen; dpInputIdx++ {
		for dpPatternIdx := 1; dpPatternIdx <= patternLen; dpPatternIdx++ {

			if pattern[dpPatternIdx-1] == '.' || pattern[dpPatternIdx-1] == input[dpInputIdx-1] {
				// when current pattern char is '.' or current pattern char == current input char:
				dp[dpInputIdx][dpPatternIdx] = dp[dpInputIdx-1][dpPatternIdx-1]
				// check '?*'
			} else if pattern[dpPatternIdx-1] == '*' {
				// ignore '?*'
				dp[dpInputIdx][dpPatternIdx] = dp[dpInputIdx][dpPatternIdx-2]
				// if the char before * matched current input char
				if pattern[dpPatternIdx-2] == '.' || pattern[dpPatternIdx-2] == input[dpInputIdx-1] {
					dp[dpInputIdx][dpPatternIdx] = dp[dpInputIdx][dpPatternIdx] || dp[dpInputIdx-1][dpPatternIdx]
				}
			}

		}
	}

	return dp[inputLen][patternLen]
}

// Go call this func in main.go
func Go() {
	common.Assert_answer(isMatch("abc", "abc"), true)
	common.Assert_answer(isMatch("abc", "a.c"), true)
	common.Assert_answer(isMatch("abc", "a*bc"), true)
	common.Assert_answer(isMatch("aaabc", "a*bc"), true)
	common.Assert_answer(isMatch("bc", "a*bc"), true)
	common.Assert_answer(isMatch("aa", "a*"), true)
}
