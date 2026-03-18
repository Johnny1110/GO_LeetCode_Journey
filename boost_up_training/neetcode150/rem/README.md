# 10. Regular Expression Matching

<br>

---

<br>

## Coding

```go
func isMatch(s string, p string) bool {
    
}
```

<br>
<br>

## Time & Space Complexity

You have a nested loop where `i` runs from `0` to `n` and `j` runs from `1` to `m`.

**Time Complexity**: Inside the loop, every operation (lookups and boolean logic) is `O(1)`. Therefore, the total work is proportional to the number of cells in your DP table: `O(n * m)`.

**Space Complexity**: You are initializing a 2D slice dp of size `(n+1) * (m+1)`. Each cell stores a boolean. This results in `O(n * m)` space.