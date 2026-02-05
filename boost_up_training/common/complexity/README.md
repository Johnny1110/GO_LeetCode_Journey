# Complexity Analysis 101

<br>

---

<br>

## What Is Complexity Analysis?

It answers two questions:

1. Time Complexity — How does runtime grow as input grows?
2. Space Complexity — How does memory usage grow as input grows?

We use **Big O** notation to express this.

| Big O Notation | Name           | Example                          |
|----------------|----------------|----------------------------------|
| O(1)           | Constant       | Accessing array by index         |
| O(log n)       | Logarithmic    | Binary search                    |
| O(n)           | Linear         | Single loop through array        |
| O(n log n)     | Linearithmic   | Efficient sorting (merge sort)   |
| O(n²)          | Quadratic      | Nested loops                     |
| O(2ⁿ)          | Exponential    | Recursive fibonacci (naive)      |

<br>

## How to Analyze: Count the Loops

Example 1: O(n)

```go
for i := 0; i < n; i++ {
    // do something O(1)
}
```

Runs `n` times → `O(n)`


<br>

Example 2: O(n²)

```go
for i := 0; i < n; i++ {
    for j := 0; j < n; j++ {
        // do something O(1)
    }
}
```

Inner loop runs `n` times, for each of n outer iterations → `O(n × n) = O(n²)`


<br>

Example 3: O(n × m)

```go
for i := 0; i < n; i++ {
    for j := 0; j < m; j++ {
        // do something O(1)
    }
}
```

Different sizes → `O(n × m)`