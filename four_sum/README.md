# 4Sum

<br>

Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:

<br>

0 <= a, b, c, d < n
a, b, c, and d are distinct.

nums[a] + nums[b] + nums[c] + nums[d] == target
You may return the answer in any order.

<br>
<br>

Example 1:
```
Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
```

Example 2:
```
Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]]
```

<br>

Constraints:
```
1 <= nums.length <= 200
-109 <= nums[i] <= 109
-109 <= target <= 109
```

<br>

## Topic 

* Array
* Two Pointers
* Sorting

<br>

## Thinking

I Think I should review 3Sum first...

<br>

Alright! After I reviewed the 3Sum and 4Sum solutions, I put together some ideas.

In the 3Sum case, we iterate through all the numbers with index i and set up two pointers: i+1 and len(nums)-1, to calculate the sum of three numbers.

In the 4Sum case, we still need a two-pointer sum approach, and we have two more numbers i and i+1.

so, it' more like:

3Sum:

```
i, i+1, len(nums)-1
```

4Sum:

```
i, i+1, i+2, len(nums)-1
```

5Sum:

```
i, i+1, i+2, i+3, len(nums)-1
```

And so on.

<br>

It's is suitable to using recursive to solve this problem.

Let's have fun there!
