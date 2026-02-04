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

I know this problem is about to using topological sort to solve. For example:

```
Input: ["hrn","hrf","er","enn","rfnn"]
```

I guess I have to create a kind of node graph, and logic is like:

1. load "hrn"

    ```go
    h -> r
    r -> n
    ```
   
2. load "hrf"

    ```
    r -> f
    ```
   
3. load "er"

   ```
    e -> r
   ```

4. load "enn"

    ```
   e -> n
   ```
   
5. load "rfnn"

   ```
   f -> n
   ```

<br>

Now we created a **Directed Graph**, head node is "h" and "e" because both of em are not be pointed by any letter.

we can using a queue (BFS) and a list to collect the result answer.

* round-1:
  * queue: ["h", "e"], result: []
  * pop "h" and push "h"'s pointer node into queue, then push "h" into result
  * queue: ["e", "r"], result: ["h"]
* round-2:
  * queue: ["e", "r"], result: ["h"]
  * pop "e" and push "e"'s pointer node into queue, then push "e" into result
  * queue: ["r", "n"], result: ["h", "e"]
* round-3:
  * queue: ["r", "n"], result: ["h", "e"]
  * pop "r" and push "r"'s pointer node into queue, then push "r" into result
  * queue: ["n", "f"], result: ["h", "e", "r"]
*  round-4:
  * queue: ["n", "f"], result: ["h", "e", "r"]
  * pop "n" and push "n"'s pointer node into queue, then push "n" into result
  * queue: ["f"], result: ["h", "e", "r", "n"]
* round-5:
  * queue: ["f"], result: ["h", "e", "r", "n"]
  * pop "f" and push "f"'s pointer node into queue, then push "f" into result
  * queue: [], result: ["h", "e", "r", "n", "f"]

result list to string = "hernf"

<br>

### Claude AI feedback

You're definitely on the right track with using topological sort! But I see a **critical misunderstanding in how you're building the graph**. 
Let me guide you through this.

<br>

#### The Key Insight You're Missing

You're extracting relationships from within each word (h→r→n from "hrn"). But that's not what the problem tells us.
Think about it this way: In a dictionary, words are sorted. What can we learn from two adjacent words being in sorted order?

For example, in English:

* "apple" comes before "banana"
* What does that tell us? Only that 'a' < 'b'
* We learn nothing about p, l, e, n within those words

<br>

#### The Correct Way to Extract Relationships

Compare adjacent words in the list, not characters within a single word.

Given your input: `["hrn", "hrf", "er", "enn", "rfnn"]`

1. Compare "hrn" vs "hrf"
   ```
   h == h  (same, keep going)
   r == r  (same, keep going)
   n != f  → we learn: n < f (n comes before f in alien alphabet)
   ```
   
2. Compare "hrf" vs "er"
   ```
   h != e  → we learn: h < e
   ```

3. Compare "er" vs "enn"
   ```
   e == e  (same, keep going)
   r != n  → we learn: r < n
   ```
   
4. Compare "enn" vs "rfnn"
   ```
   e != r  → we learn: e < r
   ```

<br>

#### Your Task

Based on this corrected understanding:

1. What edges should your directed graph actually have?
2. What would the in-degree of each node be?
3. There's also an edge case to consider: What if a longer word appears before its prefix? (e.g., ["abc", "ab"]) — What should happen?

<br>

---

<br>

## Coding

