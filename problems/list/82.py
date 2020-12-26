
# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def deleteDuplicates1(self, head):
        pre = ListNode(-1)
        pre.next = head
        p = pre
        cur = head
        while cur:
            while cur.next and cur.next.val == cur.val:
                cur = cur.next
            if pre.next == cur:
                pre = pre.next
            else:
                pre.next = cur.next
            cur = cur.next
        return p.next


if __name__ == '__main__':
    s = Solution()
    root = ListNode(1)
    root.next = ListNode(1)
    a = s.deleteDuplicates(root)
    #while a:
    #    print(a.val)