#  Letter Combinations of a Phone Number

<br>

---

<br>

# Desc

Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

![1200px-telephone-keypad2svg.png](1200px-telephone-keypad2svg.png)

<br>
<br>

Example 1:
```
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
```

Example 2:

```
Input: digits = ""
Output: []
```


Example 3:

```
Input: digits = "2"
Output: ["a","b","c"]
```

<br>

Constraints:

```
0 <= digits.length <= 4
digits[i] is a digit in the range ['2', '9'].
```

<br>
<br>

# Topic

* Hash Table
* String
* [Backtracking](https://datascientest.com/en/backtracking-what-is-it-how-do-i-use-it#:~:text=Backtracking%20is%20a%20search%20technique,optimization%2C%20planning%20and%20gaming%20problems.) 
    [Tutorial](https://medium.com/@ralph-tech/%E6%BC%94%E7%AE%97%E6%B3%95%E5%AD%B8%E7%BF%92%E7%AD%86%E8%A8%98-%E5%9B%9E%E6%BA%AF%E6%B3%95-backtracking-%E5%88%86%E6%94%AF%E5%AE%9A%E7%95%8C%E6%B3%95-branch-and-bound-29165391c377)
<br>
<br>

# Thinking

I think this is kind of ez... I'm gonna figure out what is Backtracking later and try it tomorrow (2024/07/13).

<br>

When I learn about backtrack, things become complicated.
I'm not perfect at understanding recursive solution, and backtrack is a kind of recursive.


<br>

Based on [permutation](../bonus/Fpermutation), I figure out how to utilize backtrack(DFS) to solve __Permutations__ problem.

just only need to understand what to do in first layer, and know when to stop the recursion.





