# 51. N-Queens

<br>

---

<br>

link: https://leetcode.com/problems/n-queens/description/

<br>

The n-queens puzzle is the problem of placing `n` queens on an `n x n` chessboard such that no two queens attack each other.

<br>

Given an integer `n`, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space, respectively.

<br>


Example: 

```
Input: n = 4
Output: [
    [
        ".Q..",
        "...Q", 
        "Q...",
        "..Q."
    ],
    [
        "..Q.",
        "Q...", 
        "...Q",
        ".Q.."
    ]
]
```

Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above


<br>
<br>

## Thinking

<br>

### First Shot

define `currentState` is chess board

I think this time I gonna need a `isSafeZone(currentState, x, y)` func, which can tell us is `currentState(x, y)` is in safeZone that can let us place a queen on it.


define backtracking()


```
backtracking(..., x, y): 
    return when n queens already placed.
    
    for i := x; i < n; i++:
        for j := y; h < n; n++:
            
            if !isSafeZone(currentState, i, j): 
                continue
            
            update currentState(i, j)
            backtracking(..., i, j+1)
            undo currentState(i, j)
```

<br>
<br>

### Thinking Enhancement

A queen attacks an entire row. So **we can't place two queens on the same row?**

There is `n` queens and `n` rows, which is means **one queen goes on each row**

<br>

#### One Queen Goes On Each Row

* Each recursive call handles one row
* Within that call, you only loop over columns to decide where to place the queen in that row
* Then recurse to the next row

<br>
<br>

#### Virtual Code

```
backtracking():
	if row == n -> safe currentState as one of the result

    for col := 0; col < n; col++: // iterate column
         if !isSafeZone(currentState, row, col): 
                continue
                
         update currentState with (row. col)
         goes to next row -> backtracking(..., currentState, row+1)
         undo currentState with (row. col)
```

#### How to represent `currentState`?

I think currentState could be an 2D array [][]string

<br>

#### `isSafeZone()` implement

Actually, I think `isSafeZone()` don't need row info, we just care about which column already been taken by last queen.

Like if previous 3 queens took column-1, column-2, column-3 and also they're all in safeZone, then we won't take col-1, col-2, col3 in consideration.



<br>
<br>

## Coding