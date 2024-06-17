package integer_to_roman

import (
	"go_leetcode_journey/common"
)

var romanNumValue = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var romanNumSymbol = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

func romanToInt(romain string) int {
	return 0
}

// Go call this func in main.go
func Go() {
	common.Assert_answer(romanToInt("MMMDCCXLIX"), 3749)
	common.Assert_answer(romanToInt("LVIII"), 58)
	common.Assert_answer(romanToInt("MCMXCIV"), 1994)
}
