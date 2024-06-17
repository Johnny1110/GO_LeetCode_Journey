package integer_to_roman

import (
	"go_leetcode_journey/common"
	"strings"
)

// map is not sorted, change to 2 array
var romanNumSymbolMap = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

var romanNumValue = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var romanNumSymbol = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

func intToRoman(num int) string {
	result := strings.Builder{}

	for index, value := range romanNumValue {
		for {
			temp := num - value
			if temp >= 0 {
				result.WriteString(romanNumSymbol[index])
				num = temp
			} else {
				break
			}
		}
	}
	return result.String()
}

// Go call this func in main.go
func Go() {
	common.Assert_answer(intToRoman(3749), "MMMDCCXLIX")
	common.Assert_answer(intToRoman(58), "LVIII")
	common.Assert_answer(intToRoman(1994), "MCMXCIV")
}
