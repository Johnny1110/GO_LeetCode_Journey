# Backtracking


<br>

---

<br>

## Thinking Framework: **The Decision Tree**

Backtracking is essentially exploring a decision tree. Let me guide you to discover its structure:

Question 1: If we process elements left-to-right, and at each element we decide "include it or not", what does the tree look like for [1,2]?

```
                    []
                   /  \
            include 1?  
               /        \
             [1]         []
            /   \       /   \
      include 2?     include 2?
         /   \         /    \
      [1,2]  [1]     [2]    []
```

Question 2: Where are our "answers" in this tree? (leaves? every node? specific nodes?)

<br>
<br>

## When it comes to Backtracking, Ask yourself 4 problems:

1. What represents our "current state"? (what are we building?)
2. What represents "remaining choices"? (what haven't we decided yet?)
3. When do we record a valid subset?
4. What does "undo the choice" mean concretely here?


<br>
<br>

## Important - Backtracking Pattern


###  Loop-based

```go
func backtrack(index):
record result           // every node is answer
for i := index to end:
choose nums[i]      // update state
backtrack(i + 1)    // go deeper
unchoose            // undo state
```

<br>

### Binary decision

```go
func backtrack(index):
if index == len(nums):
record result       // only leaves are answers
return

choose nums[index]      // include this element
backtrack(index + 1)
unchoose

backtrack(index + 1)    // exclude this element (no choose/unchoose needed)
```