# Longest Common Prefix

<br>

---

<br>

## Desc

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".


```
Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
```

Constraints:

```
1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters.
```

## Topic

* Trie (https://zh.wikipedia.org/zh-tw/%E5%9F%BA%E6%95%B0%E6%A0%91)

<br>

## Thinking before implements

Topic suggests using a Trie to solve this problem, so first of all, I have to implement a Trie in Golang 
(maybe it's an Object? I have no idea how to create a class in GoLang). This TrieObject looks like this:

```asciidoc
TrieObject {
  // char is the next alphabet
  Map<char, TrieObject> sub_leaves
}
```

When inputting data, fill this TrieObject first. Then iterate from the parent leaf. If the current TrieObject only has 1 sub_leaf, take the current char as part of the answer, until we find the first TrieObject with multiple sub_leaves.


<br>

## After implements

leetcode said this is a "easy" level problem,actually it's not.
I'm using a trie, but it's seems like can not solve this problem at all.
anyway, there are few approaches showed on leetcode's Editorial:

* Approach 1: [Horizontal scanning](horizontal_scanning)

* Approach 2: [Vertical Scanning](vertical_scanning)

* Approach 3: Divide and conquer

* Approach 4: Binary search (someone said it wrong, I will figure it out)

