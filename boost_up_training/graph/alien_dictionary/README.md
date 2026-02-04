# 269. Alien Dictionary

<br>

---

<br>

link: https://neetcode.io/problems/foreign-dictionary/question

<br>

## Example

```
Input: ["hrn","hrf","er","enn","rfnn"]

Output: "hernf"
```

Explanation:

* from "hrn" and "hrf", we know 'n' < 'f'
* from "hrf" and "er", we know 'h' < 'e'
* from "er" and "enn", we know get 'r' < 'n'
* from "enn" and "rfnn" we know 'e'<'r'

so one possibile solution is `"hernf"`

<br>
<br>

## Topic 

* Array
* String
* DFS, BFS
* Graph
* [Topological Sort](https://medium.com/@mswukris/%E6%8B%93%E6%92%B2%E6%8E%92%E5%BA%8F%E6%B3%95-topological-sorting-4a727d851c62)

<br>
<br>

## Thinking

### Deconstruct Problem


### Claude AI feedback


<br>

---

<br>

## Coding

