package single_number_ii

import "fmt"

func singleNumber(nums []int) int {
	counter := make([]int, 32) // int32 size

	// 1. collect counter
	for _, v := range nums {

		fmt.Printf("inspect v: %v, byte(v):%b \n", v, v)
		// each v is int32 -> 32 bit
		for i := 0; i < 32; i++ {
			mapping := 1 << i
			if mapping&v != 0 {
				counter[i] += 1
			}
		}
	}

	fmt.Printf("counter: %v \n", counter)

	result := 0
	// 2. find result in counter
	for i, v := range counter {
		if v%3 == 1 { // caught single 1
			mapping := 1 << i
			result = result | mapping
		}
	}

	return result
}
