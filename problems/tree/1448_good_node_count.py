class Solution():
    def __init__(self):
        pass

    count = 0

    def goodNodes(self, root: TreeNode) -> int:
        def dfs(node, max):
            if node is None:
                return
            if node.val >= max:
                max = node.val
                self.count += 1
            dfs(node.left, max)
            dfs(node.right, max)

        dfs(root, root.val)
        return self.count
