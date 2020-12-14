class Solution():
    def __init__(self):
        pass

    res = 0

    def findTilt(self, root: TreeNode) -> int:
        self.res = 0

        def sumTree(node):
            if node is None:
                return 0
            res = node.val
            res += sumTree(node.left)
            res += sumTree(node.right)
            return res

        def dfs(node):
            if node is None:
                return
            dfs(node.left)
            dfs(node.right)
            left = sumTree(node.left)
            right = sumTree(node.right)
            self.res += abs(left - right)

        dfs(root)
        return self.res


    def findTilt2(self, root: TreeNode) -> int:
        self.res = 0

        def dfs(node):
            if node is None:
                return
            l = dfs(node.left)
            r = dfs(node.right)
            self.res += abs(l- r)
            return node.val + l + r
        dfs(root)
        return self.res