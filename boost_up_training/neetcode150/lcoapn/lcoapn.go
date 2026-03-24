package lcoapn

var digAlphaMap = map[uint8][]uint8{
	'2': []uint8{'a', 'b', 'c'},
	'3': []uint8{'d', 'e', 'f'},
	'4': []uint8{'g', 'h', 'i'},
	'5': []uint8{'j', 'k', 'l'},
	'6': []uint8{'m', 'n', 'o'},
	'7': []uint8{'p', 'q', 'r', 's'},
	'8': []uint8{'t', 'u', 'v'},
	'9': []uint8{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	result := []string{}

	if digits == "" {
		return result
	}

	state := make([]uint8, len(digits))

	var backtracking func(idx int)
	backtracking = func(idx int) {
		if idx >= len(digits) {
			result = append(result, string(state))
			return
		}

		alphaSet := digAlphaMap[digits[idx]]

		for _, char := range alphaSet {
			// update state
			state[idx] = char
			// go next layer
			backtracking(idx + 1)
			// rollback state
			// no need
		}
	}

	backtracking(0)

	return result
}
