package multiply_strings_revamp

import (
	"go_leetcode_journey/common"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	num1Arr := splitStringToNumArray(num2)
	num2Arr := splitStringToNumArray(num1)
	res := numArrMultiply(num1Arr, num2Arr)
	return res
}

func numArrMultiply(num1 []int, num2 []int) string {
	layer1Result := [][]int{}

	for i := 0; i < len(num1)-1; i++ {

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

		layer1Result = append(layer1Result, layer2Result)
	}

	//fmt.Println(layer1Result)
	return sum(layer1Result)
}

func sum(result [][]int) string {
	galigeigei := []int{}

	resLength := len(result)
	for i := 0; i < resLength; i++ {

		buling := resLength - (i + 1)
		for x := 0; x < buling; x++ {
			result[i] = append(result[i], 0)
		}
	}

	//fmt.Println("buling:", result)

	bbLen := len(result[0])

	carryUp := 0
	for i := 0; i < bbLen; i++ {

		sumup := 0
		for _, array := range result {
			sumup += array[i]
		}

		remainer := (sumup + carryUp) % 10
		carryUp = (sumup + carryUp) / 10
		galigeigei = append(galigeigei, remainer)
	}

	//fmt.Println("galigeigei:", galigeigei)

	return makeGaligeigeiToString(galigeigei)
}

func makeGaligeigeiToString(galigeigei []int) string {
	res := ""

	skip := true
	for i := len(galigeigei) - 1; i >= 0; i-- {

		if skip && galigeigei[i] == 0 {
			continue
		}
		if skip && galigeigei[i] > 0 {
			skip = false
		}
		res += strconv.Itoa(galigeigei[i])
	}

	if res == "" {
		return "0"
	} else {
		return res
	}
}

func splitStringToNumArray(num string) []int {
	numArr := []int{}
	for i := len(num) - 1; i >= 0; i-- {
		numArr = append(numArr, int(num[i]-'0'))
	}
	numArr = append(numArr, 0)
	return numArr
}

// Go call this func in main.go
func Go() {
	common.Assert_answer(multiply("123", "456"), "56088")
	common.Assert_answer(multiply("9", "9"), "81")
	common.Assert_answer(multiply("498828660196", "840477629533"), "419254329864656431168468")
}
