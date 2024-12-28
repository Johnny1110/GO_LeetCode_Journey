# Combination Sum II

<br>

---

<br>

## Desc

Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target.

Each number in candidates may only be used once in the combination.

Note: The solution set must not contain duplicate combinations.

<br>


Example 1:

```
Input: candidates = [10,1,2,7,6,1,5], target = 8

Output:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
```


Example 2:

```
Input: candidates = [2,5,2,1,2], target = 5

Output:
[
[1,2,2],
[5]
]
```

<br>

Constraints:

```
1 <= candidates.length <= 100
1 <= candidates[i] <= 50
1 <= target <= 30
```

<br>
<br>

## Topic

* Array
* Backtracking

<br>

## Thinking

Still backtracking ! let's go.

First, create a NumQueue struct that can store numbers, calculate the sum, and easily push and pop elements.

Still following the backtracking principle, push a number into the queue, go into a deeper layer, and then roll back.

<br>

Lets having fun today! (2024//12/28)