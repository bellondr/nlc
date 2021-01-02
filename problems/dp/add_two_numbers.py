# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class AddTwoNumbers:
    def __init__(self):
        pass

    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        head = cur = ListNode()
        tag = 0
        while l1 or l2 or tag:
            cur.next = ListNode()
            if l1:
                cur.next.val = l1.val
                l1 = l1.next
            if l2:
                cur.next.val += l2.val
                l2 = l2.next
            if tag:
                cur.next.val += tag
            if cur.next.val >= 10:
                tag = 1
                cur.next.val -= 10
            else:
                tag = 0
            cur = cur.next
        return head.next