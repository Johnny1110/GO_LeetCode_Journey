package three_sum

func threeSum(nums []int) [][]int {

	// 1. sort
	quickSort(nums, 0, len(nums)-1)

	// 2. 3 pointers
	result := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ { // we need 3 element, so -2
		if nums[i] > 0 {
			break // early quick, we need a negative value here 
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue // skip duplicate (except first one)
		}

		target := -nums[i] // positive val -> what we want for 2 sums value here

		pointerA, pointerB := i+1, len(nums)-1

		for pointerA < pointerB {
			sum2 := nums[pointerA] + nums[pointerB]

			if target == sum2 {
				// found 1
				result = append(result, []int{nums[i], nums[pointerA], nums[pointerB]})
				// move 2 pointer (shrink)
				pointerA++
				for pointerA < pointerB && nums[pointerA] == nums[pointerA-1] { // try skip duplicate
					pointerA++
				}

				pointerB--
				for pointerB > pointerA && nums[pointerB] == nums[pointerB+1] {
					pointerB--
				}
			} else if target > sum2 {
				// we need a bigger val, move pointerA 
				pointerA++
				for pointerA < pointerB && nums[pointerA] == nums[pointerA-1] { // try skip duplicate
					pointerA++
				}
			} else {
				// we need a smaller val, move pointerB
				pointerB--
				for pointerB > pointerA && nums[pointerB] == nums[pointerB+1] {
					pointerB--
				}
			}
		}
	}


	return result
}

// quickSort binSearch
func quickSort(nums []int, start, end int) {

	if start >= end {
		return
	}


	cursor, writer := start, start
	
	for cursor < end {
		if nums[end] > nums[cursor] {
			nums[cursor], nums[writer] = nums[writer], nums[cursor]
			writer++
		}
		cursor++
	}

	// pivot
	nums[end], nums[writer] = nums[writer], nums[end]

	// recursive sort
	quickSort(nums, start, writer-1)
	quickSort(nums, writer+1, end)

}
