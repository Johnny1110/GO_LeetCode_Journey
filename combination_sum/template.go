package combination_sum

//
//import "fmt"
//
//type NumQueue struct {
//	cached []int
//	sum    int
//}
//
//func (queue *NumQueue) push(val int) {
//	queue.cached = append(queue.cached, val)
//	queue.sum += val
//}
//
//func (queue *NumQueue) pop() (int, bool) {
//	if len(queue.cached) == 0 {
//		return 0, false
//	}
//	val := queue.cached[len(queue.cached)-1]
//	queue.sum = queue.sum - val
//	queue.cached = queue.cached[:len(queue.cached)-1]
//	return val, true
//}
//
//func (queue *NumQueue) clone() []int {
//	result := make([]int, len(queue.cached))
//	copy(result, queue.cached)
//	return result
//}
//
//func (queue *NumQueue) clear() {
//	queue.cached = queue.cached[:0]
//	queue.sum = 0
//}
//
//func combinationSum(candidates []int, target int) [][]int {
//	results := make([][]int, 0)
//	queue := &NumQueue{[]int{}, 0}
//
//	backtracking(target, queue, 0, 0, candidates, &results)
//
//	return results
//}
//
//func backtracking(target int, queue *NumQueue, idxA, idxB int, candidates []int, results *[][]int) {
//	fmt.Println("---------------------------------------------------------")
//	if idxB >= len(candidates) {
//		idxA += 1
//		idxB = idxA
//	}
//
//	if idxA >= len(candidates) {
//		return
//	}
//
//	queue.push(candidates[idxB])
//
//	if queue.sum < target {
//		fmt.Println("lower then target:", queue.clone())
//		fmt.Println("idxA:", idxA, "idxB:", idxB)
//		backtracking(target, queue, idxA, idxB, candidates, results)
//	}
//
//	if queue.sum == target {
//
//		fmt.Println("find result!!:", queue.clone())
//		fmt.Println("idxA:", idxA, "idxB:", idxB)
//		// append results
//		*results = append(*results, queue.clone())
//		queue.pop()
//		_, b := queue.pop() // 倒退兩格，兩格失敗就退出
//		if !b {
//			return
//		}
//		backtracking(target, queue, idxA, idxB+1, candidates, results)
//	}
//
//	if queue.sum > target {
//		fmt.Println("greater then target:", queue.clone())
//		fmt.Println("idxA:", idxA, "idxB:", idxB)
//		queue.pop()
//		_, b := queue.pop() // 倒退兩格，兩格失敗就退出
//		if !b {
//			return
//		}
//		backtracking(target, queue, idxA, idxB+1, candidates, results)
//	}
//}
//
//// Go call this func in main.go
//func Go() {
//	candidates := []int{2, 3, 5}
//	res := combinationSum(candidates, 8)
//	debugResult(res)
//	// [[2,2,2,2],[2,3,3],[3,5]]
//}
//
//func debugResult(res [][]int) {
//	fmt.Println(res)
//}
//
////func backtracking(queue *NumQueue, target int, candidate []int, index int, results *[][]int) bool {
////	queue.push(candidate[index])
////
////	// 1. 當數字沒超過 target，往下繼續加相同數字
////	if queue.sum < target {
////		backtracking(queue, target, candidate, index, results)
////	}
////
////	// 2. 當數字等於 target，紀錄結果，並回傳 false
////	if queue.sum == target {
////		*results = append(*results, candidate)
////		return false
////	}
////
////	// 4. 當數字大於 target，倒退
////}
