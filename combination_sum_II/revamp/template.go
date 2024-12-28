package combination_sum_II_revamp

import (
	"fmt"
	"sort"
)

type NumQueue struct {
	container []int
	sum       int
}

func (queue *NumQueue) push(num int) {
	queue.container = append(queue.container, num)
	queue.sum += num
}

func (queue *NumQueue) pop() {
	if len(queue.container) == 0 {
		return
	}
	num := queue.container[len(queue.container)-1]
	queue.sum -= num
	queue.container = queue.container[:len(queue.container)-1]
}

func (queue *NumQueue) copy() []int {
	copyArray := make([]int, len(queue.container))
	copy(copyArray, queue.container)
	return copyArray
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates) // Sort candidates to handle duplicates easily
	result := [][]int{}
	queue := &NumQueue{}
	backtracking(candidates, target, 0, queue, &result)
	return result
}

func backtracking(candidates []int, target int, idx int, queue *NumQueue, result *[][]int) {
	if queue.sum == target {
		*result = append(*result, queue.copy())
		return
	}
	if queue.sum > target {
		return
	}

	for i := idx; i < len(candidates); i++ {
		// Skip duplicates:
		// This is the point, if i pointing at the number which is equals to previous one,
		//that means this number already try in previous step whatever succeed match or failed.
		// Skip that duplicate num is OK !
		if i > idx && candidates[i] == candidates[i-1] {
			continue
		}

		queue.push(candidates[i])
		backtracking(candidates, target, i+1, queue, result) // Move to the next index
		queue.pop()
	}
}

// Go call this func in main.go
func Go() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	res := combinationSum2(candidates, 8)
	debugResult(res)
	// ans is: [1,1,6], [1,2,5], [1,7], [2,6]
}

func debugResult(res [][]int) {
	fmt.Println(res)
}
