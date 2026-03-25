/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start

func isValid(s string) bool {
	// "([)]"
	charStack := []uint8{}
	oppositCharMap := map[uint8]uint8{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	var sPop func() (uint8, bool)
	sPop = func() (uint8, bool) {
		if len(charStack) == 0 {
			return 0, false
		}
		ret := charStack[len(charStack)-1]
		charStack = charStack[:len(charStack)-1]
		return ret, true
	}

	for _, char := range s {

		switch char {
		case '(', '[', '{':
			charStack = append(charStack, uint8(char))
		case ')', ']', '}':
			if a, ok := sPop(); ok {
				if oppositCharMap[a] != uint8(char) {
					return false
				}
			} else {
				return false
			}
		}
	}

	return len(charStack) == 0
}

// @lc code=end

