# Approach 2: Vertical Scanning

<br>

---

<br>

Based on horizontal scanning, if a very short string is at the bottom of the array, like:

```asciidoc
["aaabbbccc", "aaabbb", "aaabb", "a"]
```

Horizontal scanning still needs to loop through all the strings until it reaches the last one. 
In this case, the editorial showed me a slightly different way to solve this problem by utilizing vertical scanning.

<br>

Yeah, I did it. But there's still something I should remember:

1. In Golang or Python, "12345"[:3] means a substring from index 0 to 3-1, which results in 123.
2. When looping through characters of a string (ForLoop a string chars), if i >= len(str) means it has already reached the end,
so the loop should quit. Don't overthink it; remembering this is like following a formula. 
