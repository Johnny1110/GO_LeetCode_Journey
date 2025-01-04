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

### Data Cleaning:

Ignore irrelevant values (e.g., negatives, zeros, and numbers greater than
𝑛
n, the size of the array).

### Place Elements Where They Belong:

Use a cyclic sort or in-place placement. The idea is to rearrange the array so that `nums[i]=i+1` whenever possible.

### Identify the Gap:

After rearrangement, scan the array for the first mismatch (where `num[i] != i+1`)

The index `i+1` represents the first missing positive integer.

### Fallback Case:

If all numbers are correctly placed, the missing integer is `n+1`.


<br>
<br>

## Think it again: 

__The most important thing is the answer must greater than 0 and lower than `len(nums)`.

```
ans > 0 && ans < len(nums)
```

And the second thing we need to rearrange the input `nums[]`

we focus on the number between `0` to `len(nums)`. if we find a number which is in a wrong position, place it in right way.

<br>

For Example:

we got a int array nums: `[-50, 0, 1, 3, 5, 1000]`

the `len(num)` is 6, so we focus on the num 0~6, is this case we got :

```go
1, 3, 5
```

let's try to rearrange, put em in the right position:

`[1, 0, 3, -50, 5, 1000]`

now `[1, 3, 5]` all in the right position, `nums[0]`, `nums[2]`, `nums[4]`

<br>

then let iterate that nums with idx i in range (0 ~ 5).

if nums[i] != i+1 then answer is i+1 like:

```go
nums[i] = nums[0] = 1
1 == i+1 => pass

num[i+1] = nums[1] = 0
0 != i+1+1 => return i+1+1 = 2

final aswer is 2.
```

<br>

final answer:

```go
package main 

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		// ⬇ if number not in where it should be (nums[nums[i]-1]) .
		for nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			// Why using for ? because we swapped nums[i] to where it should be, then we got a new number in nums[i],
			// we got to check new number again. until not match up with above judgmental.
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}
```