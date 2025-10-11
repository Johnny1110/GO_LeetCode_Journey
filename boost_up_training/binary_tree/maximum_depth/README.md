# 104. Maximum Depth of Binary Tree

<br>

---


<br>

link: https://leetcode.com/problems/maximum-depth-of-binary-tree/description/

<br>

## Thinking

<br>

I remember I did something before, when I try to implement Ethereum Merkle Tree, when it comes to calculate tree height, I fund a equation for that.
but I forgot what it is...

<br>

### Asking AI to recap that equation:

Binary tree (like Merkle trees typically), the relationship between nodes and height is:Given n nodes:

```
height = ⌈log₂(n + 1)⌉ - 1
```

Or equivalently:

```
height = ⌊log₂(n)⌋
```

But LeetCode's Problem is Different!

**Key difference:** 

The LeetCode "Maximum Depth of Binary Tree" problem gives you an arbitrary binary tree that might be:

* Unbalanced (one side much deeper than the other)
* Sparse (missing nodes randomly)
* A single long chain

<br>

So nope! we can not using that formula to solve this tree height problem.

<br>
<br>

then I'm think about should I have to go through all the nodes. I guess the answer is yes.

That will be a `DFS` (Depth-First Search)

<br>
<br>

## Coding
