# 268. Missing Number 

<br>

---

<br>

link: https://leetcode.com/problems/missing-number/

<br>
<br>

Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.

```
Input: nums = [3,0,1]

Output: 2
```

<br>

## Thinking

This is a new algo concept learning about bitmap, I want to use bitmap to solve this problem.

<br>


### Decompse the problem

* The min num from input array `nums` always to be 0, max number is `len(nums)`.
* All the numbers are unique.

<br>

I think there is a one simple approach, since we already know number is `0 ~ n`, so we can calcualte the sum.

```
func sum(n int) int {
    if n is odd -> (1+n)*(n/2) + (1+n)/2

    if n - even -> (1+n)*(n/2)
}
```

then we can sum the val in `nums` like for example we know `sum(4) = 10` if there is missed 1 in nums. the sum of number is 9. 

* 10-9 = 1

1 will be the answer. 

```go
func missingNumber(nums []int) int {
	n := len(nums)
	target := sumN(n)

	numsSum := 0

	for _, v := range nums {
		numsSum += v
	}

	return target - numsSum
}

func sumN(n int) int {
	isOdd := n%2 == 1

	if isOdd {
		// is odd -> (1+n)*(n/2) + (1+n)/2
		return (1+n)*(n/2) + (1+n)/2
	} else {
		// even -> (1+n)*(n/2)
		return (1 + n) * (n / 2)
	}
}
```

Result:

![1](imgs/1.png)

<br>

But this is not what I want to do here, we are aiming for master bitmap!

<br>

One single int is 32 bit, let me list some number and see the pattern:

```
0: 0000 0000 0000 0000
1: 0000 0000 0000 0001
2: 0000 0000 0000 0010
3: 0000 0000 0000 0011
4: 0000 0000 0000 0100
5: 0000 0000 0000 0101
```

<br>

Ask yourself this: **What happens when you XOR a number with itself?**

```
5 ^ 5 = 0

     0101
XOR  0101
---------
     0000
```

<br>

* a ^ a = 0
* a ^ 0 = a

<br>

### Try This

```
nums = [3, 0, 1]    →  n = 3, so full set is [0, 1, 2, 3]

1: 0001
2: 0010
3: 0011

XOR all of 0..3:   0 ^ 1 ^ 2 ^ 3 = 0000 -> 0
XOR all of nums:   3 ^ 0 ^ 1     = 0010 -> 2
XOR those results together: 0010 -> 2
```

```
nums = [3, 0, 1, 4]    →  n = 4, so full set is [0, 1, 2, 3, 4]

1: 0001
2: 0010
3: 0011
4: 0100

XOR all of 0..4:   0 ^ 1 ^ 2 ^ 3 ^ 4 = 0100 -> 4
XOR all of nums:   3 ^ 0 ^ 1 ^ 4     = 0110 -> 6
XOR those results together: 0010 -> 2
```

<br>

That was amazing.

### XOR Properties (Fundamental)

These are the core algebraic properties of XOR:

* Identity: `a ^ 0 = a`
* Self-inverse: `a ^ a = 0`
* Commutative: `a ^ b = b ^ a`
* Associative: `(a ^ b) ^ c = a ^ (b ^ c)`

<br>

### 4-cycle pattern

```
n=0: 0                              → 0
n=1: 0 ^ 1                          → 1
n=2: 0 ^ 1 ^ 2                      → 3
n=3: 0 ^ 1 ^ 2 ^ 3                  → 0
n=4: 0 ^ 1 ^ 2 ^ 3 ^ 4              → 4
n=5: 0 ^ 1 ^ 2 ^ 3 ^ 4 ^ 5          → 1
n=6: ...                            → 7
n=7: ...                            → 0
```

<br>

| n % 4 | f(n) |
|-------|------|
| 0     | n    |
| 1     | 1    |
| 2     | n+1  |
| 3     | 0    |


<br>
<br>

## Coding

We can combined those 2 XOR in one loop.

```
iterate i from 0~n-1

    XOR i and XOR nums[i] as v

In the end, XOR the n
```

Result:

![1](imgs/1.png)