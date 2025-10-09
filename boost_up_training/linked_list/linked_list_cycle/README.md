# Linked List Cycle

<br>

---

<br>

link: https://leetcode.com/problems/linked-list-cycle/description/

<br>

>Given head, the head of a linked list, determine if the linked list has a cycle in it.

>There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. 
> Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

>Return true if there is a cycle in the linked list. Otherwise, return false.

<br>

**Constraints:**

* The number of the nodes in the list is in the range [0, 10^4]. (0 ~ ten to the fourth)
* -105 <= Node.val <= 105
* pos is -1 or a valid index in the linked-list.

<br>

## Thinking

The Constraints number-1, the number of node in the list is 0 ~ ten to the fourth, that is a big amount of elements.

I think the most important thing is how to define a linked list is cycled.

I think if every node's next is not null, that means it is a cycle.

* Using 2 pointer to target head node and current node.
* Iterate through the linked list, break if we reached same node again (which is mean cycle also).
* But if we iterate through the pointer, we may not be able to stop, because of cycle reference.
* I can only think about using O(n) 2 nested loop.

<br>
<br>

Claude AI:

**What Actually Defines a Cycle**

* A linked list has a cycle when a node's next pointer points back to a previous node in the list, creating a loop. 
    It's not just about every node having a non-null next pointer—a regular non-cyclic linked list can have that too (except for the tail, which points to null).

* The key insight: __In a cyclic list, if you keep following next pointers, you'll eventually visit the same node twice.__

* "O(n) nested loop" - This would work but is inefficient in both time O(n²) and space.

<br>

**Approach 1: Hash Set**

* Theory: Store each node you visit. If you encounter a node already in your set, there's a cycle.
* Time: O(n) - single pass
* Space: O(n) - storing all nodes
* Why it works: You're tracking history to detect revisits

<br>

**Approach 2: Floyd's Cycle Detection (Two Pointers - "Tortoise and Hare")**

This is the elegant solution worth understanding deeply:

Theory:
* Use two pointers moving at different speeds (slow moves 1 step, fast moves 2 steps)
* If there's a cycle, the fast pointer will eventually "lap" the slow pointer and they'll meet
* If there's no cycle, fast pointer reaches null

### Why does this work mathematically?

**Imagine a circular race track:**

* If both runners start at the same point but one runs twice as fast
* The faster runner will eventually catch up to the slower one from behind
* The speed difference means they close the gap by 1 position each step
* They must meet if it's circular

**In the list:**

* Once both pointers are inside the cycle, their relative distance decreases by 1 each iteration
* Since the cycle is finite, they must eventually meet
* Time: O(n) - at most 2n steps
* Space: O(1) - only two pointers!

<br>
<br>

## Coding

```java

```