
class LongestZigZagTree():
    def __init__(self):
        pass

    res = 0
    def dfs(self, node, last, length):
        if node is None:
            return
        self.res = length if length > self.res else self.res
        if last == -1:
            self.dfs(node.left, -1, 1)
            self.dfs(node.right, 1, length + 1)
        else:
            self.dfs(node.left, -1, length + 1)
            self.dfs(node.right, 1, 1)

    def longestZigZag(self, root: TreeNode) -> int:
        self.res = 0
        self.dfs(root.left, -1, 1)
        self.dfs(root.right, 1, 1)
        return self.res

