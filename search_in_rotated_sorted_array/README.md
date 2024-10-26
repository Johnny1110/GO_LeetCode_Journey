# Search in Rotated Sorted Array

<br>

---

<br>

## Desc

There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.

<br>

Example 1:

```
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
```

Example 2:

```
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
```


Example 3:
```
Input: nums = [1], target = 0
Output: -1
```

Constraints:

```
1 <= nums.length <= 5000
-104 <= nums[i] <= 104
All values of nums are unique.
nums is an ascending array that is possibly rotated.
-104 <= target <= 104

```

<br>



## Topic

* Array
* [Binary Search](https://www.bilibili.com/video/BV1Ng4y1q7E3/?spm_id_from=333.337.search-card.all.click)

<br>

## Thinking

<br>

example: [4,5,6,7,0,1,2], target = 6

1. find the pivot index, in this case is idx: 4 (which is 0)

2. we got 2 section:

    sectionA: idx 0 ~ 3 `[4,5,6,7]`

    sectionB: idx 4 ~ 6 `[0,1,2]`

3. we need to find out which section should we drill in. target 6 is greater then 4 and lower then 7, so we got sectionA

4. using Binary Search to find out where is target 6.
