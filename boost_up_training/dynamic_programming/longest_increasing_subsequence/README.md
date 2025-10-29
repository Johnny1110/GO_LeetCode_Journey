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

