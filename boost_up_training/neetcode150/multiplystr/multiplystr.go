package multiplystr

var ct int = 48

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	result := make([]int, len(num1)+len(num2))

	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {

			n1 := int(num1[i]) - ct
			n2 := int(num2[j]) - ct

			result[i+j+1] += n1 * n2
			result[i+j] += result[i+j+1] / 10
			result[i+j+1] %= 10
		}
	}

	// find first non-zero
	startIdx := 0
	for i := 0; i < len(result); i++ {
		if result[i] > 0 {
			startIdx = i
			break
		}
	}

	ans := ""
	for i := startIdx; i < len(result); i++ {
		tmp := rune(result[i] + ct)
		ans += string(tmp)
	}

	return ans
}
