# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
# Definition for singly-linked list.

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
# Definition for a binary tree node.

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution():
    def findMid(self, start, end):
        fast, slow = start, start
        while fast != end and fast.next != end:
           fast = fast.next.next
           slow = slow.next
        return slow

    def bfs(self, head, end):
        if head == end:
            return None
        mid = self.findMid(head, end)
        node = TreeNode(mid.val)
        node.left = self.bfs(head, mid)
        node.right = self.bfs(mid.next, end)
        return node

    def sortedListToBST(self, head: ListNode) -> TreeNode:
        return self.bfs(head, None)


if __name__ == '__main__':
    head = ListNode(-10, ListNode(-3, ListNode(0, ListNode(5, ListNode(9)))))
    s = Solution()
    a = s.sortedListToBST(head)
    def dfs(n):
        if n is None:
            return
        print(n.val)
        dfs(n.left)
        dfs(n.right)
    dfs(a)

