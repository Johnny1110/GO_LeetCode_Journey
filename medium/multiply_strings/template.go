package multiply_strings

import (
	"fmt"
	"go_leetcode_journey/common"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	num1Arr := splitStringToNumArray(num2)
	num2Arr := splitStringToNumArray(num1)
	intNum := numArrMultiply(num1Arr, num2Arr)
	return strconv.Itoa(intNum)
}

func numArrMultiply(num1 []int, num2 []int) int {
	layer1Result := []int{}

	for i := 0; i < len(num1); i++ {

		layer2Result := []int{}

		for b := 0; b < i; b++ {
			layer2Result = append(layer2Result, 0)
		}

		carryUp := 0
		for j := 0; j < len(num2); j++ {
			product := num1[i] * num2[j]
			product = product + carryUp
			remainder := product % 10
			carryUp = product / 10
			layer2Result = append(layer2Result, remainder)
		}

		layer1Result = append(layer1Result, reverseToInt(layer2Result))
	}

	fmt.Println(layer1Result)
	return sum(layer1Result)
}

func sum(result []int) int {
	sumup := 0
	for i := 0; i < len(result); i++ {
		sumup += result[i]
	}
	return sumup
}

func reverseToInt(result []int) int {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return arrayToInt(result)
}

func splitStringToNumArray(num string) []int {
	numArr := []int{}
	for i := len(num) - 1; i >= 0; i-- {
		numArr = append(numArr, int(num[i]-'0'))
	}
	numArr = append(numArr, 0)
	return numArr
}

func arrayToInt(arr []int) int {
	result := 0
	for _, digit := range arr {
		result = result*10 + digit
	}
	return result
}

// Go call this func in main.go
func Go() {
	common.Assert_answer(multiply("123", "456"), "56088")
	common.Assert_answer(multiply("9", "9"), "81")
}
