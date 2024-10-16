# Substring with Concatenation of All Words

<br>

---

<br>

## Desc

You are given a string s and an array of strings words. All the strings of words are of the same length.

A concatenated string is a string that exactly contains all the strings of any permutation of words concatenated.

For example, if words = ["ab","cd","ef"], then "abcdef", "abefcd", "cdabef", "cdefab", "efabcd", and "efcdab" are all concatenated strings. "acdbef" is not a concatenated string because it is not the concatenation of any permutation of words.
Return an array of the starting indices of all the concatenated substrings in s. You can return the answer in any order.


<br>

Example 1:

```
Input: s = "barfoothefoobarman", words = ["foo","bar"]

Output: [0,9]
```


Explanation:

The substring starting at 0 is "barfoo". It is the concatenation of ["bar","foo"] which is a permutation of words.
The substring starting at 9 is "foobar". It is the concatenation of ["foo","bar"] which is a permutation of words.

<br>

Example 2:

```
Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]

Output: []
```

Explanation:

There is no concatenated substring.

<br>

Example 3:

```
Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]

Output: [6,9,12]

```
Explanation:

The substring starting at 6 is "foobarthe". It is the concatenation of ["foo","bar","the"].
The substring starting at 9 is "barthefoo". It is the concatenation of ["bar","the","foo"].
The substring starting at 12 is "thefoobar". It is the concatenation of ["the","foo","bar"].

<br>
<br>

Constraints:

```
1 <= s.length <= 104
1 <= words.length <= 5000
1 <= words[i].length <= 30
s and words[i] consist of lowercase English letters.
```

<br>
<br>

## Topic

* Hash Table
* String
* [Sliding Window](https://ms0223900.medium.com/sliding-window-%E6%BC%94%E7%AE%97%E6%B3%95-%E7%B0%A1%E4%BB%8B-d967f64ac811)

<br>
<br>

## Thinking

I was thinking about using Sliding Window to iterate the input s.

I got the answer like:

1. Create a frequency map for all words in the list.

2. Slide the window across the string s, and check chunks of word-sized pieces.

3. Update the map by decrementing counts as I found valid words.

4. Check the map at the end of each window to see if all ccount are zero.

5. Handle invalid words or too many occurrences by skipping the current window when necessary.
