package mergeintervals

import "fmt"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	result := [][]int{}

	// 1 order intervals (order by first number)
	for i := 0; i < len(intervals); i++ {
		for j := i; j < len(intervals[i]); j++ {
			intA, intB := intervals[i][0], intervals[j][0]
			if intA > intB {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}

	fmt.Printf("intervals: %v\n", intervals)
	tmp := intervals[0]
	for idx := 1; idx < len(intervals); idx++ {
		if canMerge(tmp, intervals[i]) {
			tmp[1] = max(tmp[1], intervals[i][1])
		} else {
			// append result
			result = append(result, []int{tmp[0], tmp[1]})
			tmp = intervals[i]
		}
	}

	return result
}

func canMerge(intsA, intsB []int) bool {
	if intsB[0] >= intsA[0] && intsB[0] <= intsA[1] {
		return true
	}
	return false
}
