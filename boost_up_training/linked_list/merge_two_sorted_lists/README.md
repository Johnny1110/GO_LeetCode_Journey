# Merge Two Sorted Lists

<br>

----

<br>

link: https://leetcode.com/problems/merge-two-sorted-lists/description/

<br>

## Thinking

**Constraints:**

* The number of nodes in both lists is in the range [0, 50].
* -100 <= Node.val <= 100
* Both `list1` and` list2` are sorted in non-decreasing order.

<br>

Both input lists are sorted. I think that one is pretty easy, just rebuild a new linked list from both inputs.

<br>
<br>

## Coding

```java
class Solution {
    public ListNode mergeTwoLists(ListNode list1, ListNode list2) {
        if (list1 == null) {
            return list2;
        }
        if (list2 == null) {
            return list1;
        }

        ListNode header;
        // 1. choose the header
        if (list1.val < list2.val) {
            header = list1;
            list1 = list1.next;
        } else {
            header = list2;
            list2 = list2.next;
        }

        // merge the rest
        ListNode pointer = header;
        while (true) {
            if (list1 == null) {
                pointer.next = list2;
                break;
            }
            if (list2 == null) {
                pointer.next = list1;
                break;
            }

            if (list1.val < list2.val) {
                pointer.next = list1;
                pointer = list1;
                list1 = list1.next;
            } else {
                pointer.next = list2;
                pointer = list2;
                list2 = list2.next;
            }
        }

        return header;
    }
}
```

<br>

Result:

![1.png](imgs/1.png)
