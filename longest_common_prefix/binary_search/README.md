# Binary search

<br>

---

<br>

The editorial said there's a "Binary search" approach, and after I checked the solution,
I think it's not a typical binary search, at least searching data is not sorted as well.

It's more like a kinda assume the longest common prefix length.

First of all, find the min size of all the input strings, and then assume `(minLen/2)` is the result common prefix's length
compare all the string with only minLen/2, if all string matched up that mean's we got to assume `(minLen/2) + 1` next time.
vice versa, if not maching, next one will be `(minLen/2) - 1`...
