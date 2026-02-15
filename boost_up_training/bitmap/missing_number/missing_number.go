package missing_number

func missingNumber(nums []int) int {
	n := len(nums)

	sumXORN := 0
	sumXORV := 0

	for i, v := range nums {
		sumXORN ^= i
		sumXORV ^= v
	}

	sumXORN ^= n

	return sumXORN ^ sumXORV
}
