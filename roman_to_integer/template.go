package roman_to_integer

import (
	"go_leetcode_journey/common"
)

var romanSymbolVal_FULL = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

var romanSymbolVal_SIMPLE = map[uint8]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	return myIdea(s)
}

// 13 ms solved.
func myIdea(input string) int {
	totalSum := 0
	for i := len(input) - 1; i >= 0; i-- {
		currentCharVal := romanSymbolVal_SIMPLE[input[i]]
		if i-1 >= 0 {
			// calculate with left char
			leftCharVal := romanSymbolVal_SIMPLE[input[i-1]]
			if currentCharVal > leftCharVal {
				i-- // skip next
				totalSum += currentCharVal - leftCharVal
				continue
			}
		}
		totalSum += currentCharVal
	}
	return totalSum
}

// best solution, 比我聰明
func bestSol(input string) int {
	result := 0
	mapValues := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	for i := 0; i < len(input)-1; i++ {
		if mapValues[(input[i])] < mapValues[(input[i+1])] {
			result -= mapValues[(input[i])]
		} else {
			result += mapValues[(input[i])]
		}
	}
	return result + mapValues[(input[len(input)-1])]
}

// Go call this func in main.go
func Go() {
	common.Assert_answer(romanToInt("MMMDCCXLIX"), 3749)
	common.Assert_answer(romanToInt("LVIII"), 58)
	common.Assert_answer(romanToInt("MCMXCIV"), 1994)
}
