# 15. 3Sum

<br>

---

<br>

## Coding

```go
func threeSum(nums []int) [][]int {
	result := [][]int{}

	if len(nums) < 2 {
		return result
	}

	// sort
	slices.Sort(nums)

	for i := 0; i < len(nums)-2; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue // skip duplicate
		}

		if nums[i] > 0 { // we need a negative val
			return result
		}

		targetSum := -nums[i]

		// using 2 pointer
		pointerA, pointerB := i+1, len(nums)-1
		for pointerA < pointerB {

			A, B := nums[pointerA], nums[pointerB]
			tmpSum := A + B

			if tmpSum == targetSum {
				// found 1 result
				result = append(result, []int{nums[i], A, B})
				pointerA++
				for pointerA < pointerB && nums[pointerA] == nums[pointerA-1] {
					pointerA++
				}
				pointerB--
				for pointerA < pointerB && nums[pointerB] == nums[pointerB+1] {
					pointerB--
				}
			} else if tmpSum < targetSum {
				// we need a bigger num, move pointerA
				pointerA++
				for pointerA < pointerB && nums[pointerA] == nums[pointerA-1] {
					pointerA++
				}
			} else {
				// we need a smaller num, move pointerB
				pointerB--
				for pointerA < pointerB && nums[pointerB] == nums[pointerB+1] {
					pointerB--
				}
			}

		}
	}

	return result
}
```

<br>
<br>

## Time & Space Complexity

```
Assume: n = length of input nums, h = valid answer count

Time: O(log n) + O(n square) , because slices.Sort using quick sort, it is a kind of binary search operation, so it is O(log n), and 3 pointer iterate is O(n square).

Space: O(h) because we're not using extra space, except some int pointer and int tmpSum and return result is a [][]int.
```

<br>

## Gimini AI Feedback

### Time Complexity: $\mathcal{O}(n^2)$

* Sorting is actually $\mathcal{O}(n \log n)$: You mentioned slices.Sort is $\mathcal{O}(\log n)$ because it is similar to binary search. While binary search is $\mathcal{O}(\log n)$ because it cuts the search space in half each time, sorting an array requires comparing and moving all $n$ elements. Go's slices.Sort uses pattern-defeating quicksort (pdqsort), which has an average and worst-case time complexity of $\mathcal{O}(n \log n)$.

* The Loops are $\mathcal{O}(n^2)$: You have an outer loop running $n$ times, and an inner two-pointer loop that processes the remaining elements in $\mathcal{O}(n)$ time.

* Total Time: $\mathcal{O}(n \log n) + \mathcal{O}(n^2)$. In Big-O notation, we drop the less significant terms, leaving us with a final time complexity of $\mathcal{O}(n^2)$.


<br>

### Space Complexity: $\mathcal{O}(\log n)$ or $\mathcal{O}(n)$ (Auxiliary)

You estimated $\mathcal{O}(h)$ where $h$ is the number of valid answers. Whether you are right depends entirely on whether the interviewer asks for Auxiliary Space (extra space used temporarily by the algorithm) or Total Space (which includes the memory allocated for the output).

Usually, in algorithm interviews, space used to hold the output `[][]int` is not counted toward the space complexity.

* Variables: As you correctly noted, pointerA, pointerB, tmpSum, etc., only take $\mathcal{O}(1)$ constant space.

* The Hidden Cost of Sorting: This is the catch. Sorting algorithms are rarely truly zero-space. Quicksort/pdqsort requires memory for the recursion call stack. In Go, slices.Sort will take about $\mathcal{O}(\log n)$ auxiliary space.

* Conclusion: If we ignore the output array, your auxiliary space complexity is $\mathcal{O}(\log n)$. If your interviewer insists on counting the output array, then your original answer of $\mathcal{O}(h)$ (where $h$ is the number of triplets) is completely correct!