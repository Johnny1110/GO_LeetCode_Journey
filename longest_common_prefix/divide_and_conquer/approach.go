package divide_and_conquer

func DivideAndConquer(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	return longestCommonPrefix(strs, 0, len(strs)-1)

}

func longestCommonPrefix(strs []string, leftIdx int, rightIdx int) string {
	// if input is only 1 char, just return it
	if leftIdx == rightIdx {
		return strs[leftIdx]
	} else {
		// calculate the mid idx
		mid := (rightIdx + leftIdx) / 2
		// split strings by middle index until there's only 1 str
		leftPartStr := longestCommonPrefix(strs, leftIdx, mid)
		rightPartStr := longestCommonPrefix(strs, mid+1, rightIdx)
		// the final left and right will be 2 str, compare both of them find LCP
		return commonPrefix(leftPartStr, rightPartStr)
	}

}

func commonPrefix(str1 string, str2 string) string {
	// min str size
	minStrSize := min(len(str1), len(str2))
	for i := 0; i < minStrSize; i++ {
		if str1[i] != str2[i] {
			return str1[:i]
		}
	}
	return str1[:minStrSize]
}
