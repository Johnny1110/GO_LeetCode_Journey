# REDO: 3 Sum 

<br>

link: https://leetcode.com/problems/3sum/description/

<br>
<br>

## Coding - 1

```go
import "sort"

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)

	// 1. sort first
	quickSort(nums, 0, len(nums)-1)

	// 2. using 3 pointer
	for i := 0; i < len(nums)-2; i++ { // -2 because we need at least 3 values
		if nums[i] > 0 {
			continue // we only need a negative value
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue // skip duplicate
		}

		diff := -nums[i] // this is what we want for 2sum value

		pointerA, pointerB := i+1, len(nums)-1

		for pointerA < pointerB {
			sum2 := nums[pointerA] + nums[pointerB]

			if sum2 == diff { // #1. found 1 pair
				result = append(result, []int{nums[i], nums[pointerA], nums[pointerB]})

				pointerB--
				for nums[pointerB] == nums[pointerB+1] && pointerB > pointerA { // move more step if dupilcate
					pointerB--
				}

				pointerA++
				for nums[pointerA] == nums[pointerA-1] && pointerB > pointerA { // move more step if duplicate
					pointerA++
				}

			} else if sum2 > diff { //#2. we need a smaller val (move pointerB)
				pointerB--
				for nums[pointerB] == nums[pointerB+1] && pointerB > pointerA { // move more step if dupilcate
					pointerB--
				}
			} else {
				// #3. we need a bigger val (move pointerA)
				pointerA++
				for nums[pointerA] == nums[pointerA-1] && pointerB > pointerA { // move more step if duplicate
					pointerA++
				}
			}
		}

	}

	return result
}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	writer, cursor := start, start

	for cursor < end {
		if nums[end] > nums[cursor] {
			nums[writer], nums[cursor] = nums[cursor], nums[writer]
			writer++
		}
		cursor++
	}

	// swap end
	nums[end], nums[writer] = nums[writer], nums[end]

	quickSort(nums, start, writer-1)
	quickSort(nums, writer+1, end)
}
```

<br>

I'm not finish it very well, Let's do it second time.

<br>

## Coding - 2

```

```

