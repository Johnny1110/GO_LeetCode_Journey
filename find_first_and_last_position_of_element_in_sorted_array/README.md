# Find First and Last Position of Element in Sorted Array

<br>

---

<br>

Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.

<br>

Example 1:

```
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
```


Example 2:

```
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
```

Example 3:

```
Input: nums = [], target = 0
Output: [-1,-1]
```

Constraints:

```
0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums is a non-decreasing array.
-109 <= target <= 109
```

<br>
<br>

## Topic

* Array
* Binary Search

<br>

## Thinking

I'm think about using Binary Search to locate one of the target, then using two pointer moving toward to left and right side

like if the target is 7, after locate one of the 7, using 2 pointers like [1, 2, 2, 3, 4, 5, 6, 7, <(pointerA)- 7, -(pointerB)> 7, 8, 9]

when pointer A reached idx 6, and pointer B reached idx 10．return the result [pointerA+1, pointerB-1]． 

<br>

After I showed my idea to CHAT-GPT, I got a alternative method:

Here's an alternative, commonly used method:

1. Binary Search for Leftmost Occurrence: Use binary search to find the leftmost (first) occurrence of the target. You can do this by setting up the binary search, but whenever you find the target, keep moving left by adjusting the high pointer, so you ensure you get the first occurrence.

2. Binary Search for Rightmost Occurrence: Use binary search again, this time to find the rightmost (last) occurrence. When the target is found, keep moving right by adjusting the low pointer until you've located the last occurrence.







