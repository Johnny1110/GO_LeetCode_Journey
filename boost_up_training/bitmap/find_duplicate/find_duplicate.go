package find_duclicate

func findDuplicate(nums []int) int {
	slowP := 0
	fastP := 0

	// phase-1: find cicle, and locate 2 pointers at same position.
	for {
		slowP = nums[slowP]
		fastP = nums[nums[fastP]]

		if slowP == fastP {
			break
		}
	}

	// phase-2: locate cicle entrance
	fastP = 0 // reset fastP to 0
	for {
		if fastP == slowP {
			return fastP
		}

		// both move 1 step
		slowP = nums[slowP]
		fastP = nums[fastP]
	}
}
