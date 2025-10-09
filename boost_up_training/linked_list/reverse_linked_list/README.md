# 206. Reverse Linked List

<br>

link: https://leetcode.com/problems/reverse-linked-list/description/

<br>

## Thinking

That one is pretty ez. just linked list stuff. let's roll it.

<br>
<br>

## Coding

```java
class Solution {

    public class ListNode {
        int val;
        ListNode next;

        ListNode() {
        }

        ListNode(int val) {
            this.val = val;
        }

        ListNode(int val, ListNode next) {
            this.val = val;
            this.next = next;
        }
    }
    
    public ListNode reverseList(ListNode head) {
        if (head == null || head.next == null) {
            return head;
        }

        // 1. we can make sure head and head.next are not null here, so we got 2 elements at least.
        ListNode currentNode = head.next;
        ListNode previousNode = head;
        previousNode.next = null; // reset head's next to null, as it will be the tail after reverse

        while (currentNode != null) {
            // 1. store next
            ListNode nextNode = currentNode.next;
            // 2. reverse current's pointer
            currentNode.next = previousNode;
            // 3. update previousNode
            previousNode = currentNode;
            // 4. update currentNode
            currentNode = nextNode;
        }

        return previousNode;
    }
}
```

<br>

Result:

![1.png](imgs/1.png)