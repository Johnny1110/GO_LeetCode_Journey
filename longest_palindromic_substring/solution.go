package longest_palindromic_substring

import (
	"go_leetcode_journey/common"
)

func longestPalindrome(s string) string {
	//return expandAroundCenter(s)
	return dynamicProgramming(s)
}

// GO 中心擴展法 (Expand Around Center)
func expandAroundCenter(input string) string {
	// 放結果
	var bestLeftIdxODD = 0
	var bestRightIdxODD = 0

	var bestLeftIdxEVEN = 0
	var bestRightIdxEVEN = 0

	for idx, _ := range input {
		// 偶數檢查 pointers
		pointerX := idx     // 中分左
		pointerY := idx + 1 // 中分右
		// 奇數檢查 pointers
		pointerA := idx - 1 // 左邊第一個 idx
		pointerB := idx + 1 // 右邊第一個 idx

		oddNumberCheckFlag := true
		evenNumberCheckFlag := true

		for {
			if pointerA < 0 || pointerB > len(input)-1 {
				// 若超出字串邊界就結束(奇數)
				oddNumberCheckFlag = false
			}

			if pointerX < 0 || pointerY > len(input)-1 {
				// 若超出字串邊界就結束(偶數)
				evenNumberCheckFlag = false
			}

			if oddNumberCheckFlag && input[pointerA] == input[pointerB] {
				// 若左右邊一致，證明是對稱的，若比歷史最佳大，就暫存 bestLeftIdx, bestRightIdx
				if (pointerB - pointerA) > (bestRightIdxODD - bestLeftIdxODD) {
					bestLeftIdxODD = pointerA
					bestRightIdxODD = pointerB
				}
				pointerA-- // point a 往左邊走一格
				pointerB++ // point b 往右邊走一格
			} else {
				oddNumberCheckFlag = false
			}

			if evenNumberCheckFlag && input[pointerX] == input[pointerY] {
				// 若左右邊一致，證明是對稱的，若比歷史最佳大，就暫存 bestLeftIdx, bestRightIdx
				if (pointerY - pointerX) > (bestRightIdxEVEN - bestLeftIdxEVEN) {
					bestLeftIdxEVEN = pointerX
					bestRightIdxEVEN = pointerY
				}
				pointerX-- // point a 往左邊走一格
				pointerY++ // point b 往右邊走一格
			} else {
				evenNumberCheckFlag = false
			}

			if !oddNumberCheckFlag && !evenNumberCheckFlag {
				break
			}
		}
	}

	if (bestRightIdxODD - bestLeftIdxODD) > (bestRightIdxEVEN - bestLeftIdxEVEN) {
		// 奇數檢查結果比較長
		return input[bestLeftIdxODD : bestRightIdxODD+1]
	} else {
		// 偶數檢查結果比較長
		return input[bestLeftIdxEVEN : bestRightIdxEVEN+1]
	}
}

func Go() {
	common.Assert_answer_or(longestPalindrome("babad"), "bab", "aba")
	common.Assert_answer(longestPalindrome("a"), "a")
	common.Assert_answer_or(longestPalindrome("abc"), "a", "b", "c")
	common.Assert_answer(longestPalindrome("asdwGca1acGasdw"), "Gca1acG")
	common.Assert_answer(longestPalindrome("cbbd"), "bb")
	common.Assert_answer(longestPalindrome("123ababbaba123"), "ababbaba")
	common.Assert_answer(longestPalindrome("bb"), "bb")
	common.Assert_answer(longestPalindrome("abb"), "bb")
	common.Assert_answer(longestPalindrome("aaaa"), "aaaa")
}
