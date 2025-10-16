# 297. Serialize and Deserialize

<br>

---

<br>

link: https://leetcode.com/problems/serialize-and-deserialize-binary-tree/description/

<br>

> Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

> Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

> Clarification: The input/output format is the same as how LeetCode serializes a binary tree. You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.

<br>

## Thinking

I saw many example by previous tree problems, there always show me the example by an Array like:

```go
[3, 5, 1, null, null, 7, 11]
```

I'm thinking about should I use this data structure to serialize tree?

<br>
<br>

## Claude AI Suggests

There are 2 solution to implement tree serialize and deserialize

* Preorder (DFS) with null markers
* Level-order (BFS)

<br>

### Preorder (DFS) with null markers

```
     1
   /   \
  2     3
   \   /
    4 5

→ "1,2,null,4,null,null,3,5,null,null,null"
```

**Pros:** Simpler to implement recursively, natural for DFS
**Cons:** More nulls stored, looks longer

* Recursive thinking with implicit state (call stack)
* How tree structure maps to linear sequences
* The elegant symmetry between serialize/deserialize

Core skill: Understanding how recursion "automatically" handles tree traversal

<br>
<br>

### Level-order (BFS)

```
     1
   /   \
  2     3
   \   /
    4 5

→ [1, 2, 3, null, 4, 5]
```

**Pros:** Compact, intuitive visualization
**Cons:** You need to track positions carefully, handle nulls for missing children

* Explicit state management with queues
* Parent-child index relationships
* Handling "layers" of a tree

Core skill: Thinking iteratively about tree structure, managing position tracking

<br>
<br>


## Thinking - Preorder (DFS) with null markers

<br>

Let me implement DFS Preorder first.

<br>
<br>

## Coding

```go

```