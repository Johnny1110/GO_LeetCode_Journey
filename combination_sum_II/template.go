package combination_sum_II

import "fmt"

type NumQueue struct {
	container []int
	sum       int
}

func (queue *NumQueue) push(num int) {
	queue.container = append(queue.container, num)
	queue.sum += num
}

func (queue *NumQueue) pop() (bool, int) {
	if len(queue.container) == 0 {
		return false, 0
	}
	num := queue.container[len(queue.container)-1]
	queue.sum -= num
	queue.container = queue.container[:len(queue.container)-1]
	return true, num
}

func (queue *NumQueue) copy() []int {
	copyArray := make([]int, len(queue.container))
	copy(copyArray, queue.container)
	return copyArray
}

func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	queue := &NumQueue{}
	candidates = orderAsc(candidates)
	backtracking(candidates, target, 0, queue, &result)
	return result
}

func orderAsc(candidates []int) []int {
	for i := 0; i < len(candidates); i++ {
		for j := i + 1; j < len(candidates); j++ {
			if candidates[i] > candidates[j] {
				candidates[i], candidates[j] = candidates[j], candidates[i]
			}
		}
	}
	fmt.Println("candidates:", candidates)
	return candidates
}

func backtracking(candidates []int, target int, idx int, queue *NumQueue, result *[][]int) {
	// define stop point

	// stop point 1. sum is greater then target
	if queue.sum > target {
		return
	}
	// stop point 2. sum is equals to target
	if queue.sum == target {
		// check distinct
		if notDuplicate(queue.container, *result) {
			*result = append(*result, queue.copy())
		}
		return
	}

	for i := idx; i < len(candidates); i++ {
		queue.push(candidates[i])
		backtracking(candidates, target, i+1, queue, result)
		queue.pop()
	}

}

func notDuplicate(container []int, currentResult [][]int) bool {
	for _, arr := range currentResult {

		if equals(container, arr) {
			return false
		}
	}

	return true
}

func equals(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
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
