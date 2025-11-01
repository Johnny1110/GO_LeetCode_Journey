# 300. Longest Increasing Subsequence

<br>

---

<br>

link: https://leetcode.com/problems/longest-increasing-subsequence/description/

> Given an integer array nums, return the length of the longest strictly increasing subsequence.

<br>
<br>

## Thinking

<br>

It's also a DP problem.

example:

* Input: nums = `[10,9,2,5,3,7,101,18]`
* Output: `4`
* Explanation: The longest increasing subsequence is `[2,3,7,101]`, therefore the length is `4`.

<br>

### Define DP

Let's try: dp[i] = longest increasing subsequence 

<br>

Iterate through all the element, if current value is greater than previous one, `dp[i] = 1 + dp[i-1]`.

How about this? Idk, let's try this out.

* example := [10,9,2,5,3,7,101,18]
* dp := [1,1,1,2,1,2,3,1]

The result is 3, if the problem is longest consecutive increasing subsequence, this can be a correct solution.

<br>
<br>

Let's try:

dp[i] = the length of the longest increasing subsequence that __ENDS__ at index i.

<br>

Iterate through all the element, and check previous dp:

* if nums[i] > nums[x] (x is 0 ~ i-1)
* dp[i] = max(dp[i], dp[x]+1)

return the biggest value in dp as final result.


<br>
<br>

## Coding-1

```go
func lengthOfLIS(nums []int) int {
	// define DP
	dp := make([]int, len(nums))
	finalMaxVal := 1

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		currVal := nums[i]

		for j := 0; j < i; j++ {
			if currVal > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])

				if dp[i] > finalMaxVal {
					finalMaxVal = dp[i]
				}
			}
		}
	}

	return finalMaxVal
}
```

Result:

![1.png](imgs/1.png)

<br>

Of course the performance is pool. we need revamp this.

<br>
<br>

## Revamp

<br>

### The O(n log n) Optimization: Binary Search + Greedy

**The Key Insight**:

Maintain an array tails where:

* tails[i] = the smallest tail element of all increasing subsequences of length i+1

* tails[0]: Best ending value for subsequences of length 1
* tails[1]: Best ending value for subsequences of length 2


Why "smallest tail"? Because a smaller tail gives us more room to extend the subsequence.

<br>
<br>

**Step-by-Step Example**:

For `[10,9,2,5,3,7,101,18]`:

* We want to track the best ending value for each possible subsequence length.
* If we have 2 subsequences of the same length, we prefer the one with the smaller value.
* This greedy choice maximizes our chances of extending the subsequence.

<br>

1. num=10: tails = [10]
2. num=9: Replace 10 → tails = [9] (better tail for length 1)
3. num=2: Replace 9 → tails = [2] (even better tail)
4. num=5: Append → tails = [2,5] (length 2 subsequence)
5. num=3: Replace 5 → tails = [2,3] (better tail for length 2)
6. num=7: Append → tails = [2,3,7] (length 3)
7. num=101: Append → tails = [2,3,7,101] (length 4)
8. num=18: Replace 101 → tails = [2,3,7,18] (better tail for length 4)

Result: Length is 4

<br>

### Binary Search Works

* If we have a subsequence of length `i` ending with value `x`.
* And a subsequence of length `i+1` ending with value `y`.
* Then `y` must be greater than `x` (by definition of increasing subsequence)

We can use binary search to find where the current number fits.

<br>
<br>

### Coding-2

```go
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// init tails
	tails := []int{}

	// using binary-search to find tails element who is bigger than num.
	for _, num := range nums {
		idx := binarySearchIdx(&tails, num) // idx is where we can replace or append
		if idx == len(tails) {
			tails = append(tails, num)
		} else {
			tails[idx] = num
		}
	}

	return len(tails)
}

// binarySearchIdx find first element's idx bigger than num.
func binarySearchIdx(tailsPtn *[]int, num int) int {
	tails := *tailsPtn
	left, right := 0, len(tails)-1

	for left <= right {
		midIdx := (left + right) / 2
		midVal := tails[midIdx]
		if midVal > num {
			right = midIdx - 1
		} else if midVal < num {
			left = midIdx + 1
		} else {
			return midIdx
		}
	}

	return left
}
```
