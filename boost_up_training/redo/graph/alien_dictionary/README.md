# 269. Alien Dictionary

<br>

---

<br>

link: https://leetcode.com/problems/alien-dictionary

<br>

Example
Input: ["hrn","hrf","er","enn","rfnn"]

Output: "hernf"
Explanation:

from "hrn" and "hrf", we know 'n' < 'f'
from "hrf" and "er", we know 'h' < 'e'
from "er" and "enn", we know get 'r' < 'n'
from "enn" and "rfnn" we know 'e'<'r'
so one possibile solution is "hernf"

<br>

## Coding - 1

```go
func foreignDictionary(words []string) string {

}
```