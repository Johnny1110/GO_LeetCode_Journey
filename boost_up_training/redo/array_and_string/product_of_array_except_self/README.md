# 238. Product of Array Except Self

<br>

---

<br>

link: https://leetcode.com/problems/product-of-array-except-self


<br>
<br>

## Coding-1

```go

func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))
	result := make([]int, len(nums))

	pointerA, pointerB := 0, len(nums)-1

	for pointerA < len(nums) {
		if pointerA == 0 {
			// init step
			prefix[pointerA] = nums[pointerA]
			suffix[pointerB] = nums[pointerB]
		} else {
			// calculate prefix and suffix
			prefix[pointerA] = prefix[pointerA-1] * nums[pointerA]
			suffix[pointerB] = suffix[pointerB+1] * nums[pointerB]
		}

		pointerA++
		pointerB--
	}

	// collect result
	pointerA, pointerB = 0, len(nums)-1
	for pointerA <= pointerB {
		if pointerA == 0 {
			// first
			result[pointerA] = suffix[pointerA+1]
			// last
			result[pointerB] =  prefix[pointerB-1]
		} else {
			result[pointerA] = prefix[pointerA-1] * suffix[pointerA+1]
			result[pointerB] = prefix[pointerB-1] * suffix[pointerB+1]
		}

		
		pointerA++
		pointerB--
	}

	return result
}
```

This is kinda stupid, Let me revamp this.

```go
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	// 1. calculate prefix
	preProduct := 1
	for i := 0; i < len(nums); i++ {
		result[i] = preProduct
		preProduct *= nums[i]
	}

	sufProduct := 1
	for i := len(nums)-1; i >= 0; i-- {
		result[i]*= sufProduct
		sufProduct *= nums[i]
	}

	return result
}
```