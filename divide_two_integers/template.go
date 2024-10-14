package divide_two_integers

import (
	"go_leetcode_journey/common"
	"math"
)

func divide(dividend int, divisor int) int {
	// pre-process
	// if dividend is min int and divisor = -1 return max int
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	// if dividend < 0 xor divisor < 0 means result should be negative
	isNegative := (dividend < 0) != (divisor < 0)

	// abs dividend and divisor
	absDividend := abs(dividend)
	absDivisor := abs(divisor)

	// Initialize quotient
	quotient := 0

	for absDividend >= absDivisor {
		tempDivisor := absDivisor
		multiple := 1
		times := 0

		for absDividend >= (tempDivisor << 1) {
			tempDivisor <<= 1
			multiple <<= 1
			times++
		}

		absDividend = absDividend - (absDivisor << times)

		quotient += multiple

	}

	if isNegative {
		return -quotient
	} else {
		return quotient
	}
}

func divide2(dividend int, divisor int) int {
	// pre-process
	// if dividend is min int and divisor = -1 return max int
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	// if dividend < 0 xor divisor < 0 means result should be negative
	isNegative := (dividend < 0) != (divisor < 0)

	// abs dividend and divisor
	absDividend := abs(dividend)
	absDivisor := abs(divisor)

	// Initialize quotient
	quotient := 0

	// in this loop. try to decrease absDividend until absDividend < absDivisor
	for absDividend >= absDivisor {
		tempDivisor := absDivisor
		multiple := 1

		// tempDivisor << 1 : tempDivisor multiply 2
		for absDividend >= (tempDivisor << 1) {
			tempDivisor <<= 1
			multiple <<= 1
		}

		absDividend -= tempDivisor
		quotient += multiple
	}

	if isNegative {
		quotient = -quotient
	}

	return quotient
}

func abs(dividend int) int {
	if dividend < 0 {
		return -dividend
	} else {
		return dividend
	}
}

func Go() {
	common.Assert_answer(divide2(100, 3), 33)
	common.Assert_answer(divide2(10, 3), 3)
	common.Assert_answer(divide2(-10, 3), -3)
	common.Assert_answer(divide2(10, -3), -3)
	common.Assert_answer(divide2(1000, 10), 100)
	common.Assert_answer(divide2(100, 100), 1)
	common.Assert_answer(divide2(100, 101), 0)
}
