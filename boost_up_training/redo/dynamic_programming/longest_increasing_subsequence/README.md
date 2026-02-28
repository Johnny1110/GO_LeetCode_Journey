# 300. Longest Increasing Subsequence 

<br>

---

<br>

link: https://leetcode.com/problems/longest-increasing-subsequence

<br>
<br>

```
Input: nums = [10,9,2,5,3,7,101,18]

Output: 4

Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
```

<br>

## Coding-1

```go

func lengthOfLIS(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	bestAns := 1

	// dp lengthOfLIS on idx
	dp := make([]int, len(nums))
	dp[0] = 1

	for idx := 1; idx < len(nums); idx++ {

		tmpLIS := 1
		currentNum := nums[idx]

		for j := 0; j < idx; j++ {
			if nums[j] < currentNum {
				tmpLIS = max(dp[j]+1, tmpLIS)
			}
		}

		dp[idx] = tmpLIS
		bestAns = max(bestAns, tmpLIS)

	}

	return bestAns
}
```

Too Slow: 34ms

<br>

## Coding-2 revamp

```go

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	// better tail
	tails := []int{}
	tails = append(tails, nums[0])

	for idx := 1; idx < len(nums); idx++ {
		currentNum := nums[idx]

		targetTailPosition := binarySearch(tails, currentNum)

		if targetTailPosition == len(tails) {
			tails = append(tails, currentNum)
		} else {
			tails[targetTailPosition] = currentNum
		}

	}

	return len(tails)
}

func binarySearch(tails []int, num int) int {
	startIdx, endIdx := 0, len(tails)-1

	for startIdx <= endIdx {

		midIdx := (startIdx + endIdx) / 2

		if tails[midIdx] == num {
			return midIdx
		} else if tails[midIdx] < num {
			startIdx = midIdx + 1
		} else {
			endIdx = midIdx - 1
		}

	}

	return startIdx
}
```