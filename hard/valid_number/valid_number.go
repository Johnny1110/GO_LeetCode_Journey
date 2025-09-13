package valid_number

// all the following are valid numbers: "2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789",
// while the following are not valid numbers: "abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53".
func isNumber(s string) bool {
	for _, v := range s {
		if !(v >= 48 && v <= 57) {
			return false
		}
	}

	return true
}
