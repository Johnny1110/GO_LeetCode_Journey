package multiply_strings_revamp_2

import (
	"go_leetcode_journey/common"
	"strings"
)

func multiply(num1 string, num2 string) string {

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	num1Len, num2Len := len(num1), len(num2)
	// 1. create a result array (len = num1Len + numLen2)
	result := make([]uint8, num1Len+num2Len)

	// 2. iterate through the num1 and num2
	for i := num1Len - 1; i >= 0; i-- {
		for j := num2Len - 1; j >= 0; j-- {
			product := (num1[i] - '0') * (num2[j] - '0')
			temp := product + result[i+j+1]
			result[i+j+1] = temp % 10
			result[i+j] += temp / 10
		}
	}

	// Convert result array to string (skip leading zeros)
	var sb strings.Builder
	skip := true
	for _, digit := range result {
		if digit == 0 && skip {
			continue
		}
		skip = false
		sb.WriteByte(byte(digit) + '0') // Convert int to byte
	}

	return sb.String()
}

// Go call this func in main.go
func Go() {
	common.Assert_answer(multiply("123", "456"), "56088")
	common.Assert_answer(multiply("9", "9"), "81")
	common.Assert_answer(multiply("498828660196", "840477629533"), "419254329864656431168468")
}
