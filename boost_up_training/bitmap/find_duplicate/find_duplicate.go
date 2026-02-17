package find_duclicate

func findDuplicate(nums []int) int {
	// 1 <= n <= 100000
	bitmap := make([]byte, 100000/8)

	for _, v := range nums {
		v -= 1
		bitmapIdx := (v / 8)
		bitIdx := v % 8

		byteGroup := bitmap[bitmapIdx]
		// using AND

		mapping := 1 << bitIdx

		if int(mapping)&int(byteGroup) > 0 {
			return v + 1
		} else {
			bitmap[bitmapIdx] = byteGroup | byte(mapping)
		}
	}

	return 0
}
