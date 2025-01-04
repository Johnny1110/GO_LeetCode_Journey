# First Missing Positive

<br>

---

<br>

## Desc

Given an unsorted integer array nums. Return the smallest positive integer that is not present in nums.

You must implement an algorithm that runs in O(n) time and uses O(1) auxiliary space.

<br>

Example 1:

```
Input: nums = [1,2,0]
Output: 3
```

Explanation: The numbers in the range [1,2] are all in the array.

<br>

Example 2:

```
Input: nums = [3,4,-1,1]
Output: 2
```

Explanation: 1 is in the array but 2 is missing.

<br>


Example 3:

```
Input: nums = [7,8,9,11,12]
Output: 1
```

Explanation: The smallest positive integer 1 is missing.

<br>

Constraints:

```
1 <= nums.length <= 105
-231 <= nums[i] <= 231 - 1
```

<br>

## Topic

* Array
* Hash Table

<br>
<br>

## Thinking

<br>

I just implement it roughly:

```go
func firstMissingPositive(nums []int) int {
	nums = append(nums, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		expectNext := nums[i] + 1
		if expectNext < nums[i+1] && expectNext > 0 {
			return expectNext
		}
	}
	return nums[len(nums)-1] + 1
}
```

<br>

I know it not a good approach to solve this problem. let me ask CHAT-GPT for a help.

### Problem-Solving Process

* Understand the Problem: What are the inputs and expected outputs? Are there any constraints (e.g., time/space complexity)?

* Explore Examples: Work through small examples to find patterns or properties.

* Plan a Strategy: Use the patterns identified to outline a logical approach before coding.

* Evaluate Efficiency: Check if your solution meets the constraints and optimize if needed.

<br>





