# Search in Rotated Sorted Array (Refactor)

<br>

---

<br>

1. Pivot Finding Logic: The pivot-finding loop can be made clearer. You should consider returning immediately if you find the target during pivot detection.

2. Binary Search: The binSearch function can be simplified and made iterative to avoid stack overhead from recursion.

3. Handling Edge Cases: Ensure that edge cases like arrays with a single element or duplicates are handled appropriately.
