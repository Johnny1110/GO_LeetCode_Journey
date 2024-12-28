package combination_sum_perfect_ans

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
	results := make([][]int, 0)
	queue := &NumQueue{}
	backtracking(target, queue, 0, candidates, &results)
	return results
}

func backtracking(target int, queue *NumQueue, idx int, candidates []int, results *[][]int) {

	if queue.sum == target {
		*results = append(*results, queue.clone())
		return
	}

	if queue.sum > target {
		return
	}

	for i := idx; i < len(candidates); i++ {
		queue.push(candidates[i])
		backtracking(target, queue, i, candidates, results)
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
