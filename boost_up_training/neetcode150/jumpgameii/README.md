# 45. Jump Game II

<br>

---

<br>

```
Input: nums = [2,3,1,1,4]
Output: 2
```

Explanation: The minimum number of jumps to reach the last index is 2. 
Jump 1 step from index 0 to 1, then 3 steps to the last index.

<br>
<br>

## Coding - DFS (TLE)

```go
func jump(nums []int) int {
	endIdx := len(nums) - 1

	var dfs func(idx, ans int) int
	dfs = func(idx, ans int) int {
		ans += 1

		if idx >= endIdx {
			return ans
		}

		tmp := math.MaxInt32
		for step := 1; step <= nums[idx]; step++ {
			tmp = min(tmp, dfs(idx+step, ans))
		}

		return tmp
	}

	return dfs(0, 0) - 1
}
```

Time Limit Exceeded (TLE)

<br>

## Coding - DFS Revamp with DP

```go
func jump(nums []int) int {
	n := len(nums)

	// init memory
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(idx int) int
	dfs = func(idx int) int {
		if idx >= n-1 {
			return 0
		}

		if memo[idx] != -1 {
			return memo[idx]
		}

		minJumps := math.MaxInt32

		maxJump := nums[idx]
		for step := 1; step <= maxJump; step++ {
			// The jumps needed is 1 (the current jump) + the min jumps from the next index
			minJumps = min(minJumps, 1+dfs(idx+step))
		}

		memo[idx] = minJumps
		return minJumps
	}

	return dfs(0)
}
```

<br>
<br>

### Jump Game II is famous for its Greedy solution, which runs in $O(N)$ time.

Instead of exploring paths recursively, think about the array in terms of "ranges" or "levels" (similar to a Breadth-First Search).

1. From index 0, what is the farthest index you can reach in exactly 1 jump? That range is your first "level."

2. As you iterate through that first level, keep track of the absolute farthest index you could reach if you took a second jump.

3. The moment your iteration reaches the end of your current level, you are forced to make another jump. You then update your current level's boundary to be the farthest index you discovered.

<br>

## Coding - Greedy

```go
func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	farest := 0
	jumps := 0
	currentEnd := 0

	for i := 0; i < len(nums); i++ {
		step := nums[i]
		farest = max(farest, i+step)

		if i == currentEnd {
			jumps++
			currentEnd = farest
		}

		if currentEnd >= len(nums)-1 {
			break
		}
	}

	return jumps
}
```

<br>
<br>

## Time & Space Complexity

```
Assume: n = input length of nums

Time: O(n)

Space: O(1)
```