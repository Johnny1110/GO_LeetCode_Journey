package insert_interval

func insert(intervals [][]int, newInterval []int) [][]int {

	finalResult := [][]int{}
	if len(intervals) == 0 {
		return append(finalResult, newInterval)
	}

	mergedInterval := []int{newInterval[0], newInterval[1]}
	flag := true

	for _, interval := range intervals {
		if mergedInterval[0] > interval[1] { // situation-1: No-Overlapping (new interval both right)
			finalResult = append(finalResult, interval)
		} else if mergedInterval[1] < interval[0] { // situation-2: No-Overlapping (new interval both left)
			if flag {
				flag = false
				finalResult = append(finalResult, mergedInterval)
			}
			finalResult = append(finalResult, interval)
		} else { // situation-3: Overlapping (do merge)
			// Merge the intervals by expanding newInterval's boundaries
			mergedInterval[0] = min(interval[0], mergedInterval[0])
			mergedInterval[1] = max(interval[1], mergedInterval[1])
		}
	}

	if flag {
		finalResult = append(finalResult, mergedInterval)
	}

	return finalResult
}

func Go() {}
