# 3 Sum

<br>

---

<br>

A man's dream.

<br>

## Desc

Given an integer array nums, return all the triplets `[nums[i], nums[j], nums[k]]` 

such that `i != j, i != k, and j != k`,

and `nums[i] + nums[j] + nums[k] == 0`.

<br>

Notice that the solution set must not contain duplicate triplets.


```
Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
```

<br>

```
Example 2:

Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
Example 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.
```

<br>

Constraints:

```
3 <= nums.length <= 3000
-105 <= nums[i] <= 105
```

<br>

## Topic

* Array
* Two Pointers
* Sorting

<br>

## Thinking

1. seems like I got to sort that input array then utilized two pointer.

## with ChatGPT hint

First of all, sort all the integers from the input nums array. Then, iterate through the nums array with an index i.

Make 2 pointers:

```
left: i+1 
right: len(muns)-1
```

Sum num[i], num[left], num[right], if sum result = 0, add those triplets to the answer list.
then increment left by 1, and decrement right by 1.

if num[i], num[left], num[right] sum result is less then 0, move the left index forward (to the right)
Otherwise, move right index backward (to the left)

and remember skip duplicates value for the iterate.

<br>

