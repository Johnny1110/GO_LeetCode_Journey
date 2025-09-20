# 238. Product of Array Except Self

<br>

---

<br>

link: https://leetcode.com/problems/product-of-array-except-self/description/


Given an integer array nums, return an array answer such that answer[i] is equal to the product (乘積) of all the elements of nums except nums[i].

<br>

runs in O(n) time and without using the division operation.

<br>
<br>

## Thinking

<br>

### The Topic:

* Array
* Prefix Sum


### What is Prefix Sum?

Prefix Sum is a powerful preprocessing technique that allows us to answer range sum queries in constant time after an initial O(n) setup. 
It's based on the mathematical concept of cumulative sums.

```go
// Given an array arr = [a₀, a₁, a₂, ..., aₙ₋₁], the prefix sum array prefix is:
prefix[0] = arr[0]
prefix[1] = arr[0] + arr[1]  
prefix[2] = arr[0] + arr[1] + arr[2]
...
prefix[i] = arr[0] + arr[1] + ... + arr[i]
```

<br>

Now I have figure out how to utilize prefix sum...

<br>

I have no idea, I need some hint:

* Think how you can efficiently utilize prefix and suffix products to calculate the product of all elements except self for each index. Can you pre-compute the prefix and suffix products in linear time to avoid redundant calculations?


* Can you minimize additional space usage by reusing memory or modifying the input array to store intermediate results?

<br>

Damn.. that one is really hard.

I saw the hint, prefix and suffix, this is critical clue.

let me do a example:

```go
nums = [1,2,3,4]

// the product prefix is:
prefix = [1, 2, 6, 24]
// the product suffix is:
suffix = [24 ,24, 12, 4]

// final result
result = [24, 12, 8, 6]
```

let zoom in the result;
```
result[0] = prefix[-1] * suffix[1] = 1 * 24 = 24
result[1] = prefix[0] * suffix[2] = 1 * 12 = 12
result[2] = prefix[1] * suffix[3] = 2 * 4 = 8
result[3] = prefix[2] * surfix[4] = 6 * 1 = 6 
```

<br>

I got it, let's try it out.

<br>
<br>

## Coding-1

```go
func productExceptSelf(nums []int) []int {
	totalLen := len(nums)

	// 1. calculate product prefix.
	prefix := make([]int, totalLen)
	prefix[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] * nums[i]
	}

	// 2. calculate product suffix.
	suffix := make([]int, totalLen)
	suffix[totalLen-1] = nums[totalLen-1]
	for i := totalLen - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i]
	}

	// 3.calculate final result.
	finalResult := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {

		pre := 1
		suf := 1

		if i > 0 {
			pre = prefix[i-1]
		}
		if i < totalLen-1 {
			suf = suffix[i+1]
		}

		finalResult[i] = pre * suf
	}

	return finalResult
}
```

<br>

![1.png](imgs/1.png)

<br>

Good, I'm already figure out the point. Now let's refine the detail.

<br>

## Refine

Check out the hint, minimize additional space usage by reusing memory or modifying the input array to store intermediate results


<br>

I removed the `finalResult` slice, and modifying the input array instead.

<br>

```go
func productExceptSelf(nums []int) []int {
	totalLen := len(nums)

	// 1. calculate product prefix.
	prefix := make([]int, totalLen)
	prefix[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] * nums[i]
	}

	// 2. calculate product suffix.
	suffix := make([]int, totalLen)
	suffix[totalLen-1] = nums[totalLen-1]
	for i := totalLen - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i]
	}

	// 3.calculate final result.
	//finalResult := make([]int, len(nums)) // remove this
	for i := 0; i < len(nums); i++ {

		pre := 1
		suf := 1

		if i > 0 {
			pre = prefix[i-1]
		}
		if i < totalLen-1 {
			suf = suffix[i+1]
		}
		nums[i] = pre * suf // using nums instead.
	}
	return nums
}
```

<br>

![2.png](imgs/2.png)

<br>

Good! let's see what we can improve here further.

<br>

```go
func productExceptSelf(nums []int) []int {
	totalLen := len(nums)

	// 1. calculate product prefix.
	prefix := make([]int, totalLen)
	suffix := make([]int, totalLen)

	prefix[0] = nums[0]
	suffix[totalLen-1] = nums[totalLen-1]

	pointerA := 1
	pointerB := totalLen - 2
	// optimized: using 1 loop
	for pointerA < totalLen {
		prefix[pointerA] = prefix[pointerA-1] * nums[pointerA]
		suffix[pointerB] = suffix[pointerB+1] * nums[pointerB]
		pointerA++
		pointerB--
	}

	// 3.calculate final result.
	for i := 0; i < len(nums); i++ {

		pre := 1
		suf := 1

		if i > 0 {
			pre = prefix[i-1]
		}
		if i < totalLen-1 {
			suf = suffix[i+1]
		}

		nums[i] = pre * suf
	}

	return nums
}
```

I optimized with only using 1 loop to calculate prefix and suffix. but the performance is not improved.

![3.png](imgs/3.png)

Let's ask for some help.

### Claude AI

Your algorithm is theoretically optimal at O(n) time, but there are several constant-factor optimizations we can make:

### Issues with Current Implementation:

* Extra space: Using two separate arrays (prefix and suffix)
* Multiple passes: Three separate loops
* Memory allocation overhead: Two make() calls
* Cache misses: Accessing different arrays in different patterns

<br>

### The Key Optimization: Single Array with Two Passes

The insight is that we can reuse the output array to store intermediate results, eliminating the need for separate prefix/suffix arrays.

<br>

### Optimized Approach Theory:

* Pass 1: Build prefix products in the result array
    ```
    result[i] = product of all elements to the LEFT of i
    ```

* Pass 2: Multiply by suffix products on-the-fly

    ```
    result[i] *= product of all elements to the RIGHT of i
    ```

<br>
<br>

## Coding-Final

<br>

```go
func productExceptSelf(nums []int) []int {
	totalLen := len(nums)

	result := make([]int, totalLen)
	result[0] = 1 // no left element for index 0

	// Pass-1: Fill in the prefix:
	for i := 1; i < totalLen; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	// Pass-2: Fill in the suffix:
	temp := 1
	for i := totalLen - 1; i >= 0; i-- {
		result[i] *= temp
		temp *= nums[i]
	}

	return result
}
```

<br>

![4.png](imgs/4.png)
