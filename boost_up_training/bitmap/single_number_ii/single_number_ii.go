package single_number_ii

func singleNumber(nums []int) int {
	counter := make([]int, 32) // int32 size

	// 1. collect counter
	for _, v := range nums {
		// each v is int32 -> 32 bit
		for i := 0; i < 32; i++ {
			mapping := 1 << i
			if mapping&v != 0 {
				counter[i] += 1
			}
		}
	}

	var result int32
	result = 0
	// 2. find result in counter
	for i, v := range counter {
		if v%3 == 1 { // caught single 1
			mapping := 1 << i
			result = result | int32(mapping)
		}
	}

	return int(result)
}
