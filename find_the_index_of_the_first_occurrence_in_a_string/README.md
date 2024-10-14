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



