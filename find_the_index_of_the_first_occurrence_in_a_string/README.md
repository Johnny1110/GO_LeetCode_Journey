# Find the Index of the First Occurrence in a String

<br>

---

<br>

# Desc

Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

<br>

Example 1:
```
Input: haystack = "sadbutsad", needle = "sad"
Output: 0
```
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.

<br>

Example 2:

```
Input: haystack = "leetcode", needle = "leeto"
Output: -1
```

Explanation: "leeto" did not occur in "leetcode", so we return -1.

<br>

Constraints:
```
1 <= haystack.length, needle.length <= 104
haystack and needle consist of only lowercase English characters.
```

<br>
<br>

## Topic

* Two Pointers
* String
* String Matching

<br>
<br>

## Thinking

I think just using 2 pointers A and B.
Pointer B iterate through the string. if B reach the first char of needle, move Pointer A to Pointer B's position directly.
If all matched up return Pointer A as the answer.

My solution is:

```go
func strStr(haystack string, needle string) int {
    if len(needle) > len(haystack) {
        return -1
    }

    needleIdx := 0
    pointerA := 0
    for pointerB := 0; pointerB < len(haystack); pointerB++ {
        if haystack[pointerB] == needle[needleIdx] {
            needleIdx++
            if needleIdx == len(needle) {
                return pointerA
            }
        } else {
            pointerA++
            pointerB = pointerA - 1
            needleIdx = 0
        }`
    }
    return -1
}
```

<br>

After that, I asked ChatGPT for helping me improve my answer, and here's the ChaGPT answer:

```
Your solution using two pointers is a reasonable approach to solve this problem, but it's not the most efficient or cleanest. The logic is sound, but resetting pointerB to pointerA - 1 after a mismatch adds extra complexity. This can be optimized. Here's some analysis and an alternative approach.

Issues with Current Approach:
Pointer Reset Overhead: After a mismatch, you're resetting pointerB to pointerA - 1, which forces the loop to restart checking characters that have already been checked. This can lead to unnecessary iterations.
Comparisons: You're manually managing when to move the pointerA and needleIdx. There are cleaner ways to handle this without multiple variables.
Optimized Approach (Sliding Window):
You can simplify the solution by leveraging Go's substring comparison. Here's an alternative using a sliding window approach, which directly compares substrings of the same length as needle within haystack.

Alternative Solution:
```

```go
func strStr(haystack string, needle string) int {
    n, m := len(haystack), len(needle)

    if m > n {
        return -1
    }

    for i := 0; i <= n-m; i++ {
        if haystack[i:i+m] == needle {
            return i
        }
    }

    return -1
}
```

```
Explanation:

The loop runs from i = 0 to n - m, where n is the length of haystack and m is the length of needle.
At each iteration, you check the substring of length m starting at index i to see if it matches needle.
If a match is found, you return the index i. Otherwise, after the loop ends, you return -1.
```


