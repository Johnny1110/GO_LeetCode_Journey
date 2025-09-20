package product_of_array_except_self

func productExceptSelf(nums []int) []int {
	totalLen := len(nums)

	result := make([]int, totalLen)
	result[0] = 1 // no left element for index 0

	// Pass-1: Fill in the prefix:
	for i := 1; i < totalLen; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	// Pass-2: Fill in the suffix:
	temp := 1
	for i := totalLen - 1; i >= 0; i-- {
		result[i] *= temp
		temp *= nums[i]
	}

	return result
}
