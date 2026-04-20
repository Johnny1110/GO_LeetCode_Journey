/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {
	targetIdx := len(nums) - 1

	queue := Queue([]int{})
	idxCursor := 0
	farestReach := 0
	queue.Push(0)

	for queue.Len() != 0 {

		levelLen := queue.Len()
		for range levelLen {
			thisIdx := queue.Pop()
			steps := nums[thisIdx]
			reachable := thisIdx + steps

			// early break
			if reachable >= targetIdx {
				return true
			}

			farestReach = max(farestReach, reachable)
		}

		// add some new idx into queue for next round.
		for i := idxCursor + 1; i <= farestReach; i++ {
			queue.Push(i)
		}
		// update idxCursor
		idxCursor = farestReach
	}

	return false
}

type Queue []int

func (q Queue) Len() int {
	return len(q)
}

func (q *Queue) Push(val int) {
	*q = append(*q, val)
}

func (q *Queue) Pop() int {
	ret := (*q)[0]
	(*q) = (*q)[1:]
	return ret
}

// @lc code=end

