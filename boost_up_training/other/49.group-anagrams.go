/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	// Use [26]byte to drastically reduce hashing overhead
	resMap := make(map[[26]byte][]string)

	for _, str := range strs {
		var count [26]byte

		// Inline the counting logic to avoid function call overhead
		for i := 0; i < len(str); i++ {
			count[str[i]-'a']++
		}

		resMap[count] = append(resMap[count], str)
	}

	res := make([][]string, 0, len(resMap))
	for _, v := range resMap {
		res = append(res, v)
	}

	return res
}

// @lc code=end

