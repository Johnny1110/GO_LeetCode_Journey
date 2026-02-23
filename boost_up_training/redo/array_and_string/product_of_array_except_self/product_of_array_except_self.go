package product_of_array_except_self

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	// 1. calculate prefix
	preProduct := 1
	for i := 0; i < len(nums); i++ {
		result[i] = preProduct
		preProduct *= nums[i]
	}

	sufProduct := 1
	for i := len(nums)-1; i >= 0; i-- {
		result[i]*= sufProduct
		sufProduct *= nums[i]
	}

	return result
}
