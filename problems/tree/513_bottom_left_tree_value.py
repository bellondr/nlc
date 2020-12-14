class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class FindBottomLeftTreeValue():
    def __init__(self):
        pass

    res = 0
    maxLevel = 0

    # 深度
    def dfs(self, node, length):
        if node is None:
            return

        if self.maxLevel < length + 1:
            self.maxLevel = length + 1
            self.res = node.val

        self.dfs(node.left, length + 1)
        self.dfs(node.right, length + 1)

    # 广度
    def bfs(self, node, q, index):
        if node is None:
            return
        if len(q) <= index:
            a = [node.val]
            q.append(a)
        else:
            q[index].append(node.val)
        self.bfs(node.left, q, index+1)
        self.bfs(node.right, q, index+1)


    def findBottomLeftValue(self, root):
        q = []
        self.bfs(root, q, 0)
        return q[len(q)-1][0]


if __name__ == '__main__':
    root = TreeNode(1, TreeNode(2, TreeNode(3), TreeNode(4)), TreeNode(9))
    s = FindBottomLeftTreeValue()
    a =  s.findBottomLeftValue(root)
    print(a)