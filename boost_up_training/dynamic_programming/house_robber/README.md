# 198. House Robber

<br>

----

<br>

link: https://leetcode.com/problems/house-robber/description/

<br>
<br>

## Thinking


<br>
<br>

## Coding

```go
func rob(nums []int) int {
	dp := make([]int, len(nums))
	if len(nums) >= 1 {
		dp[0] = nums[0]
	}
	if len(nums) >= 2 {
		dp[1] = max(nums[0], nums[1])
	}
	if len(nums) >= 3 {
		dp[2] = max(nums[0]+nums[2], nums[1])
	}

	for i := 3; i < len(nums); i++ {

		a := dp[i-3] + nums[i]
		b := dp[i-2] + nums[i]
		c := dp[i-1]

		dp[i] = max(a, b, c)
	}

	return dp[len(dp)-1]
}
```

<br>

result:

![1.png](imgs/1.png)

<br>

I passed the unit test, but I still feel like kinda weird.

## Revamp

```go
a := dp[i-3] + nums[i] // Rob house i and best from 3 houses ago
b := dp[i-2] + nums[i] // // Rob house i and best from 2 houses ago  
c := dp[i-1] // I don't rob this house
```

* dp[i-2] already contains the maximum of all possibilities up to house i-2
* This includes dp[i-3] (it was one of the options when calculating dp[i-2])
* So comparing dp[i-3] + nums[i] is redundant!

### The Core Pattern (Theory)

The House Robber problem has a beautiful, simple recurrence relation:
"At each house, you have exactly TWO choices:"

* Rob this house: Take nums[i] + dp[i-2] (current house + best from 2 houses ago)
* Skip this house: Take dp[i-1] (best from previous house)

That's it! Just two choices, not three.

<br>

```go
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	// init dp
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		// 1. profit from rob 2 house behind and rob this house
		profitA := dp[i-2] + nums[i]
		// 2, profit from rob previous house and skip this house
		profitB := dp[i-1]
		dp[i] = max(profitA, profitB)
	}

	return dp[len(nums)-1]
}
```
