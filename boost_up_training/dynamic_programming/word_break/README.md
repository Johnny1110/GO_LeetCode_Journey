# 139. Word Break

<br>

---

<br>

link: https://leetcode.com/problems/word-break/description/

<br>
<br>

## Thinking

**What are you trying to determine?**

* Think about what information you need at each position in the string
* What does it mean for a substring to be "valid" in this context?

**What smaller problems could help solve the bigger problem?**

* If you know something about the first part of the string, how does that help with the rest?
* Consider: if you can break `s[0...i]` into valid words, what do you need to check for `s[0...j]` where `j > i`?

<br>

### Hint for DP State:

Think about this: 

* Instead of trying to break the entire string at once, what if you asked yourself a simpler question for each position in the string?

* The DP state often answers a yes/no question about a prefix of your string. What would be a useful yes/no question to ask about each prefix?

<br>
<br>


## Coding 

```go
func wordBreak(s string, wordDict []string) bool {
    
}
```