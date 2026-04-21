/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// sort the intervals by first int element O(log2 n)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// merge
	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]

		if last[1] >= intervals[i][0] { // can merge
			// merge
			last[1] = max(last[1], intervals[i][1])
		} else {
			// add new result
			result = append(result, intervals[i])
		}
	}

	return result
}

// Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].

// Input: intervals = [[4,7],[1,4]]
// Output: [[1,7]]
// Explanation: Intervals [1,4] and [4,7] are considered overlapping.
// @lc code=end

