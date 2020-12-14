# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def widthOfBinaryTree(self, root: TreeNode) -> int:
        self.m = {}
        self.max = -1
        def dfs(node, level, num):
            if node is None:
                return
            dfs(node.left, level + 1, num * 2 + 1)
            dfs(node.right, level + 1, num * 2)
            if level not in self.m:
                self.m[level] = [num]
            else:
                self.m[level].append(num)
        dfs(root, 1, 0)
        for v in self.m:
            self.m[v].sort()
            self.max = max(self.max, self.m[v][-1] - self.m[v][0])
        return self.max + 1

    def widthOfBinaryTree2(self, root: TreeNode) -> int:
        if root is None:
            return 0
        q = deque([(root, 0)])
        res = 0
        while q:
            res = max(res, q[-1][1] - q[0][-1] + 1)
            for i in range (len(q)):
                node, index = q.popleft()
                if node.left:
                    q.append((node.left, index * 2 ))
                if node.right:
                    q.append((node.right, index *2 + 1))
        return res