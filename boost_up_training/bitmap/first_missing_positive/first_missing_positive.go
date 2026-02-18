package first_missing_positive

func firstMissingPositive(nums []int) int {
	numsLen := len(nums)

	// stage-1. sort the nums
	for i := 0; i < numsLen; i++ {
		v := nums[i]

		if v <= 0 || v > numsLen {
			// discard: mark as -1
			nums[i] = -1
		} else {
			// do swap v to where it should be (j).
			j := v - 1

			if i == j {
				continue // already sitting in right place
			}

			nums[i], nums[j] = nums[j], nums[i]
			if nums[i] == nums[j] { // duplicate case
				nums[i] = -1 // discard nums[i]
			}
			i-- // sort nums[i] again, until it be discarded.
		}

	}

	// stage-2, find min index which val is not negative
	for i, v := range nums {
		if v != i+1 {
			return i + 1
		}
	}

	return numsLen + 1
}
