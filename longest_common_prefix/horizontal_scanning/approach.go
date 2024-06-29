package horizontal_scanning

import "strings"

func HorizontalScanning(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 { // loop when prefix is not "" or could not match strs[i]
			prefix = prefix[:len(prefix)-1] // prefix remove 1 char from the right side
			if prefix == "" {               // if not matching unit the end then fk it up
				return ""
			}
		}
	}
	return prefix
}
