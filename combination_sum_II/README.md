# Combination Sum II

<br>

---

<br>

## Desc

Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target.

Each number in candidates may only be used once in the combination.

Note: The solution set must not contain duplicate combinations.

<br>


Example 1:

```
Input: candidates = [10,1,2,7,6,1,5], target = 8

Output:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
```


Example 2:

```
Input: candidates = [2,5,2,1,2], target = 5

Output:
[
[1,2,2],
[5]
]
```

<br>

Constraints:

```
1 <= candidates.length <= 100
1 <= candidates[i] <= 50
1 <= target <= 30
```

<br>
<br>

## Topic

* Array
* Backtracking

<br>

## Thinking

Still backtracking ! let's go.

First, create a NumQueue struct that can store numbers, calculate the sum, and easily push and pop elements.

Still following the backtracking principle, push a number into the queue, go into a deeper layer, and then roll back.

<br>

Lets having fun today! (2024//12/28)

I actually solved that problem with backtracking, but I don't know how to distinct duplicate array in a proper way.

I roughly write a solution like this:

```go
package main

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
```

Before append array in to result, check duplicate. I knew it's a little bit stupid, but I already try my best.

<br>
<br>

## Let's ask Professor Chat-GPT for a help

<br>

Chat-GPT's feedback:

Your approach is on the right track, but there are more efficient and "pro" ways to handle duplicates without manually checking each combination in the result. 
Instead, you can leverage sorting and skip duplicate candidates during the recursive process. Here's an improved version of your solution:

__Key Changes__:

* Sorting and Skipping Duplicates in the Loop: By sorting the candidates array first, duplicate numbers will appear consecutively. During backtracking, you can skip duplicates by comparing the current candidate with the previous one.


* No Need for notDuplicate: The skipping mechanism during recursion ensures no duplicate combinations are added to the result.

<br>

Here’s the improved solution:

```go
package main

import "sort"

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

	for i := idx; i < len(candidates); i++ {
		// Skip duplicates
		if i > idx && candidates[i] == candidates[i-1] {
			continue
		}

		// Stop early if the current number exceeds the remaining target
		if queue.sum+candidates[i] > target {
			return
		}

		queue.push(candidates[i])
		backtracking(candidates, target, i+1, queue, result) // Move to the next index
		queue.pop()
	}
}
```

<br>

### Explanation of Improvements:

* Sorted Input: Sorting ensures duplicates are adjacent, making them easier to handle.

* Skipping Duplicates: if i > idx && candidates[i] == candidates[i-1] skips processing the same number at the same recursive depth, preventing duplicate combinations.

* Early Stopping: if queue.sum + candidates[i] > target prevents unnecessary recursive calls when the target cannot be reached.


### Why This is Better:

* Efficiency: Avoids the notDuplicate check, which compares each result combination.

* Clean Logic: The skipping mechanism makes handling duplicates straightforward.

* Optimal: Early stopping avoids exploring invalid branches.

* Feel free to try this and let me know if you have any questions!

<br>

Alright! This content is quite interesting, let's digest it.