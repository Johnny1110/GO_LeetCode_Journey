# 46. Permutations

<br>

---

<br>

link: https://leetcode.com/problems/permutations/description/

<br>

Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

<br>

Example 1:

```
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
```

Example 2:

```
Input: nums = [0,1]
Output: [[0,1],[1,0]]
```

Example 3:

```
Input: nums = [1]
Output: [[1]]
```

<br>
<br>

## Thinking

<br>

Ask myself 4 problem:

1. What represents our "current state"? (what are we building?)
2. What represents "remaining choices"? (what haven't we decided yet?)
3. When do we record a valid subset?
4. What does "undo the choice" mean concretely here?

<br>

* Answer-1: currentState is a `[]int` which is store our choice of input `nums`
* Answer-2: like [1 ,2, 3], if we choice 1, there is [2, 3] left as remaining choices
* Answer-3: I think we can record a valid subset when nothing left in remaining choices 
* Answer-4: after deep diving (after recursive call to next level), currentState = currentState[: len(currentState)-1]


<br>

### The Backtracking Pattern

```
func backtracking:

    if currentState's length is full -> record currentState into result collection.
    for x:= range := unused
        update currentState currentState = append(currentState, x)
        deep diving to next decision level
        undo, reset x -> currentState = currentState[:len(currentState)-1]
```

<br>

### How do we implements unused?



<br>
<br>

## Coding