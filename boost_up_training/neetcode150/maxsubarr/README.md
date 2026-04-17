# 53. Maximum Subarray

<br>

---

<br>

## Example:

```
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.
```

## Coding 

```go
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxV := nums[0]
	dp := nums[0]

	for i := 1; i < len(nums); i++ {
		prevSum := dp
		curr := nums[i]

		if prevSum < 0 {
			dp = curr
		} else {
			dp = prevSum + curr
		}

		maxV = max(maxV, dp)
	}

	return maxV
}
```

<br>
<br>

## Time & Space Complexity

```
Assume: n = input nums length.

Time: O(n) - 1 iteration.

Space: O(1) - instead of dp []int, here I using a dp int.
```