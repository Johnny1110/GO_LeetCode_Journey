package find_duclicate

import (
	"math"
)

func findDuplicate(nums []int) int {
	// 1 <= n <= 100000
	bitmap := make([]byte, 100000/8)

	for _, v := range nums {
		bitmapIdx := (v / 8)
		bitIdx := v % 8

		if bitIdx == 0 {
			bitmapIdx -= 1
		}

		byteGroup := bitmap[bitmapIdx]

		mapping := math.Pow(2, float64(bitIdx)) // mapping value could be [1, 2, 4, 8, 16, 32, 64, 128]

		// using AND
		// bitIdx will be 0~7, so we push it to 1~8 with adding 1
		bitIdx += 1

		if int(mapping)&int(byteGroup) > 0 {
			return v
		} else {
			bitmap[bitmapIdx] = byteGroup | byte(mapping)
		}
	}

	return 0
}
