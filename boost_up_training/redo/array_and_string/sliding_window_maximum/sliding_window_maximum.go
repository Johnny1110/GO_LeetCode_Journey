package sliding_window_maximum

// MonoQueue contains nums's index, biggest one always at the left
type MonoQueue []int

func (q MonoQueue) Len() int {
	return len(q)
}

func (q *MonoQueue) PushRight(idx int) {
	*q = append(*q, idx)
}

func (q *MonoQueue) PopLeft() (int, bool) {
	if q.Len() == 0 {
		return -1, false
	}

	idx := (*q)[0]
	*q = (*q)[1:] // remove front

	return idx, true
}

func (q *MonoQueue) PopRight() (int, bool) {
	if q.Len() == 0 {
		return -1, false
	}

	idx := (*q)[q.Len()-1]
	*q = (*q)[:q.Len()-1]

	return idx, true
}

func (q MonoQueue) PeekLeft() (int, bool) {
	if q.Len() == 0 {
		return -1, false
	}

	return q[0], true
}

func (q MonoQueue) PeekRight() (int, bool) {
	if q.Len() == 0 {
		return -1, false
	}

	return q[q.Len()-1], true
}

// ========================================================

func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0)

	// init MonoQueue
	Q := MonoQueue(make([]int, 0))

	// init first window
	for idx := 0; idx < k; idx++ {
		if Q.Len() == 0 {
			// nothing in queue, just push
			Q.PushRight(idx)
		} else {
			// check left
			for {
				if rightIdx, exists := Q.PeekRight(); exists {
					if nums[rightIdx] > nums[idx] {
						Q.PushRight(idx)
						break
					} else {
						Q.PopRight()
					}
				} else {
					Q.PushRight(idx)
					break
				}
			}
		}
	}

	// collect frist result
	leftIdx, _ := Q.PeekLeft()
	result = append(result, nums[leftIdx])

	// start sliding windows
	for idx := k; idx < len(nums); idx++ {
		// 1. check left, expire old idx
		if leftIdx, exists := Q.PeekLeft(); exists {
			if leftIdx+k-1 < idx { // expired
				Q.PopLeft()
			}
		}

		// 2. check right, remove small val
		for {
			if rightIdx, exists := Q.PeekRight(); exists {
				if nums[rightIdx] > nums[idx] {
					Q.PushRight(idx)
					break
				} else {
					Q.PopRight()
				}
			} else {
				Q.PushRight(idx)
				break
			}
		}

		// 3. collect result
		leftIdx, _ := Q.PeekLeft()
		result = append(result, nums[leftIdx])

	}

	return result
}
