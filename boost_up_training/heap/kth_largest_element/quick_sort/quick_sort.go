package quick_sort

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	start, end := 0, len(nums)-1
	sortPartition(nums, start, end)

}

func sortPartition(nums []int, startIdx, endIdx int) {
	pivot := endIdx

	if startIdx == endIdx {
		return
	}

	i, j := startIdx, startIdx // init 2 pointers

	// i: where the next "small" element should be placed
	// j: scans through the array

	for j < pivot {

		if nums[j] < nums[pivot] {
			// swap i&j
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}

		j++
	}

	// Swap nums[i] with pivot at final
	nums[pivot], nums[i] = nums[i], nums[pivot]

	pivot = i

	if pivot > startIdx {
		sortPartition(nums, startIdx, pivot-1)
	}

	if pivot < endIdx {
		sortPartition(nums, pivot+1, endIdx)
	}
}
