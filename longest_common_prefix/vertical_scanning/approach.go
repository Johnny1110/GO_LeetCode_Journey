package vertical_scanning

func VerticalScanning(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	firstStr := strs[0]                  // set first string as the std
	for i := 0; i < len(firstStr); i++ { // loop the first string's every char

		for j := 1; j < len(strs); j++ { // loop the other strs

			if i >= len(strs[j]) || firstStr[i] != strs[j][i] { // if not matching or str len not enough
				return firstStr[:i] // substring (0 ~ i-1)
			}
		}
	}

	return firstStr
}
