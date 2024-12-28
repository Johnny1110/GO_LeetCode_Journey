package combination_sum_self_try

import "fmt"

type NumQueue struct {
	cached []int
	sum    int
}

func (queue *NumQueue) push(val int) {
	queue.cached = append(queue.cached, val)
	queue.sum += val
}

func (queue *NumQueue) pop() (int, bool) {
	if len(queue.cached) == 0 {
		return 0, false
	}
	val := queue.cached[len(queue.cached)-1]
	queue.sum = queue.sum - val
	queue.cached = queue.cached[:len(queue.cached)-1]
	return val, true
}

func (queue *NumQueue) clone() []int {
	result := make([]int, len(queue.cached))
	copy(result, queue.cached)
	return result
}

func (queue *NumQueue) clear() {
	queue.cached = queue.cached[:0]
	queue.sum = 0
}

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	theQueue := &NumQueue{}
	backtracking(candidates, target, 0, theQueue, &result)
	return result
}

func backtracking(candidates []int, target int, idx int, queue *NumQueue, result *[][]int) {

	// stop point 1. if total sum is grater then target -> just return
	if queue.sum > target {
		return
	}
	// stop point 2. if total sum is equals to target -> add to result then return
	if queue.sum == target {
		*result = append(*result, queue.clone())
	}

	for i := idx; i < len(candidates); i++ {
		// backtracking standard flow (add one element and go next layer then rollback)
		queue.push(candidates[i])
		backtracking(candidates, target, i, queue, result)
		queue.pop()
	}
}

// Go call this func in main.go
func Go() {
	candidates := []int{2, 3, 5}
	res := combinationSum(candidates, 8)
	debugResult(res)
	// [[2,2,2,2],[2,3,3],[3,5]]
}

func debugResult(res [][]int) {
	fmt.Println(res)
}
