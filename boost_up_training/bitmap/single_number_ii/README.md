# 137. Single Number II

<br>

---

<br>

link: https://leetcode.com/problems/single-number-ii

<br>

Given an integer array nums where every element appears three times except for one, which appears exactly once. Find the single element and return it.

You must implement a solution with a linear runtime complexity and use only constant extra space.


<br>

Example 1:

```
Input: nums = [2,2,3,2]
Output: 3
```


Example 2:

```
Input: nums = [0,1,0,1,0,1,99]
Output: 99
```

<br>
<br>

## Thinking

In previous Single Number I problem, every element appears twice except for one, only appear once.

Last time we used `XOR` to solve the problem, because `a ^ a = 0`

<br>

But this time, it is definatly not about XOR.

<br>

### The insight: think about each bit position independently.

If every number appears exactly 3 times except one, then for any given bit position, the total count of 1s across all numbers is either:

* A multiple of 3 (from the tripled numbers)
* A multiple of 3 plus 1 (the single number has a 1 at that bit)


```
nums = [2, 2, 3, 2]


In binary:
2 = 10
2 = 10
3 = 11
2 = 10

bit 0 count: 1  →  1 % 3 = 1  →  answer bit 0 = 1
bit 1 count: 4  →  4 % 3 = 1  →  answer bit 1 = 1

answer = 11 = 3 

```

Since the problem constraints typically have values fitting in 32-bit integers, you only ever have 32 bit positions. So a simple [32]int array is enough:

```
index:  0  1  2  3  4  ... 31
count:  0  0  0  0  0  ... 0
```

<br>
<br>

## Coding


```go
func singleNumber(nums []int) int {
    
}
```