I /*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	
	charFreqMapA := make(map[byte]int)
	for _, char := range s1 {
		charFreqMapA[byte(char)]++
	}

	windowSize := len(s1) // window size
	// ------------------------------------------
	
	// init window
	charFreqMapB := make(map[byte]int)
	for i := 0; i < windowSize; i++ {
		char := s2[i]
		charFreqMapB[byte(char)]++
	}

	ptnA, ptnB := 0, windowSize-1

	for ptnB < len(s2) {
		// 1. compare 2 map
		if same(charFreqMapA, charFreqMapB) {
			return true
		}
		
		
		// 2. move both ptn forward
		charFreqMapB[s2[ptnA]]--
		ptnA++
		ptnB++
		if (ptnB >= len(s2)) {
			return false // out of window
		}
		charFreqMapB[s2[ptnB]]++
		
		
	}

	return false
}
// "ab"
// "eidboaoo"
func same(mapA, mapB map[byte]int) bool {
	// mapA is from s1

	for k, v := range mapA {
		if mapB[k] != v {
			return false
		}
	}

	return true
}
// Input: s1 = "ab", s2 = "eidbaooo"
// @lc code=end

