# Approach 1: Horizontal Scanning

<br>

---

<br>

Referencing the editorial, it seems to iterate through all the strings,
comparing two of them in each loop to find the longest common prefix.
It loops until the end and then returns the longest result.

I'm on it.

<br>

After half an hour, I realized I didn't fully understand the problem. 
I thought "common prefix" meant that any two strings could match up. 
However, the truth is that each string should match the same prefix, 
and that will be the answer.



like:

```asciidoc
["flow", "ass", "fly"] ans => ""
["flow", "fl", "floooo"] ans => "fl"
```


<br>

let's try again.
